(ns anagram
  (:use [clojure.string :only [lower-case]]))

(defn anagram? [word potential-anagram] 
  (and
    (not= (lower-case word) (lower-case potential-anagram))
    (= (sort (lower-case word)) (sort (lower-case potential-anagram)))))

(defn anagrams-for [word, potential-anagrams]
  (filter (partial anagram? word) potential-anagrams))