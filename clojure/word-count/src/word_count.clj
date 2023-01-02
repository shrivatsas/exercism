(ns word-count)

(defn get-words [sentence]
  (re-seq #"\w+" (clojure.string/lower-case sentence)))

(defn map-word-count [dict word]
  (assoc dict word (inc (dict word 0))))

(defn word-count [sentence]
  (reduce map-word-count {} (get-words sentence)))