;; to write a function first write +
;; what is fibonacci series f# , clojure (because immutable by default     )
;; ->> -> macro works like a pipeline-> threading micro
;; BIG 3 map/filter/reduce use in the data pipelines
;; let creates local binding 
;; def is used for global 
;; in functional no mutable state 
;; no value can be misaligned
;; it reads like a description of records
;; #(...) for a placeholder (+ )
;; inc -> increase every number by one
;; assoc - > adds or changes a key , returns a new map
;; 2 types of functions pure and impure 
;;  atom is the managed mutable state
;; impure function is made with "!" -> can acheived by mutable funtion
;; pure function will never compromise with safety deterministic(they will return the same )-> it alwasys uses the value supplied by the function it is not globally claimed 
;; in clojure ir is immutable if we have to add in list it will ifrst create a new vector and show it
;; pure vs impure imp
;; task1:create a map and a key value pair then print original vs updated
;; task2:define a pure function with named discout and take two arguments with price and pct print price after discount 
;; 
;; threading macro 
;; if we use macro -> we start with the input 
;; thrread-last macro(->>) imp (how to use thread-last macro)
;; (->> [1 2 3 4 5 6]
;;      (filter even?)
;;      (map #(* % %))
;;      (reduce +))
;; 
;; reduce -> reduces whole list into single value
;; do twice 
;; map filter e=reduce we use thread -last
;; try and catch expression
;; safe division funtion with try and catch expression 
;; (def result              //binding ->def
;;   (try
;;      (/ 10 2)
;;      (catch ArithematicException e
;;        (println "Error:" (.getMessage e))
;;        0)))
;; result
;; 
;; 
;; testing definition
;;(.method obj arg) 
;;(Math/sqrt x)
;;(Date.)


