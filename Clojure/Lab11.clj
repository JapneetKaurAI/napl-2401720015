;; 
;; (import '(java.util ArrayList HashMap))
;; 
;; ;; 1. Calling Java Static Methods
;; (println "Absolute Value:" (Math/abs -25))
;; (println "Square Root:" (Math/sqrt 64))
;; (println "Current Time:" (System/currentTimeMillis))
;; 
;; ;; 2. Creating Java Object Instances
;; (def arr (new ArrayList))
;; (def mp (new HashMap))
;; 
;; ;; 3. Invoking Instance Methods
;; (.add arr "Apple")
;; (.add arr "Banana")
;; (.add arr "Mango")
;; 
;; (.put mp "Name" "Japneet")
;; (.put mp "Age" 20)
;; 
;; (println "Java ArrayList:" arr)
;; (println "Java HashMap:" mp)
;; (println "List Size:" (.size arr))
;; (println "Map Name:" (.get mp "Name"))
;; 
;; ;; 4. Converting Clojure and Java Collections
;; (def clj-list [10 20 30 40])
;; 
;; (def java-list (java.util.ArrayList. clj-list))
;; 
;; (println "Java List:" java-list)
;; 
;; (def converted-list (vec java-list))
;; 
;; (println "Clojure Vector:" converted-list)
;; 



;; Function to throw a Custom Exception
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
