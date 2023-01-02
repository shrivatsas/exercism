(ns bob
  (:use [clojure.string :only [blank? trim upper-case ends-with?]]))

(defn all-uppercase? [s]
  (= s (upper-case s)))

(defn alphabets? [s]
  (string? (re-find #"[A-Z]+" s)))

(defn response-for [input]
  (cond
    (blank? input) "Fine. Be that way!"
    (and
      (all-uppercase? input)
      (alphabets? input)) "Whoa, chill out!"
    (ends-with? input "?") "Sure."
    :else "Whatever."
  )
)