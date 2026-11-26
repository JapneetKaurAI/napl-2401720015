(defn check-age [age]
  (if (< age 18)
    (throw (ex-info "Age must be 18 or above"
                    {:age age :type :invalid-age}))
    (println "Eligible to vote")))

;; 1. Handling ArithmeticException
(try
  (println "Division:" (/ 10 0))
  (catch ArithmeticException e
    (println "Arithmetic Error:" (.getMessage e)))
  (finally
    (println "Arithmetic block completed")))

;; 2. Handling NullPointerException
(try
  (let [x nil]
    (.length x))
  (catch NullPointerException e
    (println "Null Pointer Error:" (.getMessage e)))
  (finally
    (println "Null block completed")))

;; 3. Handling Custom Exception
(try
  (check-age 16)
  (catch clojure.lang.ExceptionInfo e
    (println "Custom Error:" (ex-message e))
    (println "Error Details:" (ex-data e))
    (println "Graceful recovery: Please enter a valid age."))
  (finally
    (println "Custom exception block completed")))