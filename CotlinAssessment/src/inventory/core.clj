
(ns inventory.core
  (:require [ring.adapter.jetty :refer [run-jetty]]
            [ring.middleware.json :refer [wrap-json-body wrap-json-response]])
  (:import [java.sql DriverManager]))

;; DATABASE CONNECTION
(def db-url "jdbc:sqlite:inventory.db")


;; 1. CREATE DATABASE AND TABLE

(defn create-table []
  (try
    (Class/forName "org.sqlite.JDBC")

    (with-open [conn (DriverManager/getConnection db-url)
                stmt (.createStatement conn)]

      (.execute stmt
                "CREATE TABLE IF NOT EXISTS items (
           id INTEGER PRIMARY KEY,
           name TEXT NOT NULL,
           quantity INTEGER NOT NULL,
           price REAL NOT NULL
         )"))

    (println "Database connected successfully!")

    (catch Exception e
      (println "Database Error:" (.getMessage e))
      (throw e))))


;; 2. ADD ITEM TO DATABASE

(defn add-item! [id name quantity price]
  (try
    ;; Input validation
    (when (or (not (integer? id))
              (not (pos? id))
              (not (string? name))
              (clojure.string/blank? name)
              (not (integer? quantity))
              (neg? quantity)
              (not (number? price))
              (neg? price))
      (throw (ex-info "Invalid item details" {})))

    (with-open [conn (DriverManager/getConnection db-url)
                stmt (.prepareStatement conn
                                        "INSERT INTO items (id, name, quantity, price)
                   VALUES (?, ?, ?, ?)")]

      (.setInt stmt 1 id)
      (.setString stmt 2 name)
      (.setInt stmt 3 quantity)
      (.setDouble stmt 4 (double price))

      (.executeUpdate stmt))

    {:message "Item added successfully!"}

    (catch Exception e
      {:error (.getMessage e)})))


;; 3. READ ITEMS FROM DATABASE

(defn get-items []
  (try
    (with-open [conn (DriverManager/getConnection db-url)
                stmt (.createStatement conn)
                rs (.executeQuery stmt "SELECT * FROM items")]

      (loop [result []]
        (if (.next rs)
          (recur
           (conj result
                 {:id (.getInt rs "id")
                  :name (.getString rs "name")
                  :quantity (.getInt rs "quantity")
                  :price (.getDouble rs "price")}))
          result)))

    (catch Exception e
      (println "Read Error:" (.getMessage e))
      (throw e))))


;; 4. UPDATE ITEM

(defn update-item! [id quantity]
  (try
    (when (or (not (integer? id))
              (not (integer? quantity))
              (neg? quantity))
      (throw (ex-info "Invalid ID or quantity" {})))

    (with-open [conn (DriverManager/getConnection db-url)
                stmt (.prepareStatement conn
                                        "UPDATE items SET quantity = ? WHERE id = ?")]

      (.setInt stmt 1 quantity)
      (.setInt stmt 2 id)

      (if (pos? (.executeUpdate stmt))
        {:message "Item updated successfully!"}
        {:error "Item not found"}))

    (catch Exception e
      {:error (.getMessage e)})))


;; 5. DELETE ITEM

(defn delete-item! [id]
  (try
    (with-open [conn (DriverManager/getConnection db-url)
                stmt (.prepareStatement conn
                                        "DELETE FROM items WHERE id = ?")]

      (.setInt stmt 1 id)

      (if (pos? (.executeUpdate stmt))
        {:message "Item deleted successfully!"}
        {:error "Item not found"}))

    (catch Exception e
      {:error (.getMessage e)})))


;; 6. DATA PIPELINE USING MAP, FILTER AND REDUCE

(defn process-data [items]

  ;; MAP: Calculate the value of each item
  (let [mapped-items
        (map (fn [item]
               (assoc item :total-value
                      (* (:quantity item) (:price item))))
             items)

        ;; FILTER: Select items with available quantity
        filtered-items
        (filter (fn [item]
                  (pos? (:quantity item)))
                mapped-items)

        ;; REDUCE: Calculate total inventory value
        summary
        (reduce
         (fn [result item]
           (-> result
               (update :item-count inc)
               (update :total-value
                       + (:total-value item))))
         {:item-count 0 :total-value 0.0}
         filtered-items)]

    {:items (vec filtered-items)
     :summary summary}))


;; 7. RING HTTP HANDLER

(defn handler [request]
  (try
    (let [uri (:uri request)
          method (:request-method request)]

      (cond

        ;; GET ALL ITEMS
        (and (= uri "/items") (= method :get))
        {:status 200
         :body {:items (get-items)}}

        ;; GET PROCESSED SUMMARY
        (and (= uri "/summary") (= method :get))
        {:status 200
         :body (process-data (get-items))}

        ;; ADD NEW ITEM
        (and (= uri "/items") (= method :post))
        (let [body (:body request)
              result (add-item!
                      (:id body)
                      (:name body)
                      (:quantity body)
                      (:price body))]
          {:status (if (:error result) 400 201)
           :body result})

        ;; UPDATE ITEM QUANTITY
        (and (= uri "/items") (= method :put))
        (let [body (:body request)
              result (update-item!
                      (:id body)
                      (:quantity body))]
          {:status (if (:error result) 400 200)
           :body result})

        ;; DELETE ITEM
        (and (= uri "/items") (= method :delete))
        (let [body (:body request)
              result (delete-item! (:id body))]
          {:status (if (:error result) 404 200)
           :body result})

        ;; UNKNOWN ROUTE
        :else
        {:status 404
         :body {:error "Endpoint not found"}}))

    (catch Exception e
      {:status 500
       :body {:error (.getMessage e)}})))


;; 8. START SERVER

(defn -main []
  (create-table)

  (println "Starting Clojure Inventory Server...")
  (println "Server running at http://localhost:3000")
  (println "GET    /items")
  (println "GET    /summary")
  (println "POST   /items")
  (println "PUT    /items")
  (println "DELETE /items")

  (run-jetty
   (-> handler
       (wrap-json-body {:keywords? true})
       (wrap-json-response))
   {:port 3000
    :join? true}))
