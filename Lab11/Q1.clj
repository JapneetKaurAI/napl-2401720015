
(import '(java.util ArrayList HashMap))

;; 1. Calling Java Static Methods
(println "Absolute Value:" (Math/abs -25))
(println "Square Root:" (Math/sqrt 64))
(println "Current Time:" (System/currentTimeMillis))

;; 2. Creating Java Object Instances
(def arr (new ArrayList))
(def mp (new HashMap))

;; 3. Invoking Instance Methods
(.add arr "Apple")
(.add arr "Banana")
(.add arr "Mango")

(.put mp "Name" "Japneet")
(.put mp "Age" 20)

(println "Java ArrayList:" arr)
(println "Java HashMap:" mp)
(println "List Size:" (.size arr))
(println "Map Name:" (.get mp "Name"))

;; 4. Converting Clojure and Java Collections
(def clj-list [10 20 30 40])

(def java-list (java.util.ArrayList. clj-list))

(println "Java List:" java-list)

(def converted-list (vec java-list))

(println "Clojure Vector:" converted-list)

