(ns clock
  (:require [clojure.string :refer [join]]))

(defn clock->string [clock]
  (join ":" (map #(format "%02d" %) clock))
)

; get absolute timestamp value
(defn abs [timestamp]
  (mod timestamp (* 24 60))
)

; convert to a timestamp value
(defn timestamp [hours minutes]
  (+ (* hours 60) minutes)
)

(defn clock [hours minutes]
  (let [timestamp (abs (timestamp hours minutes))]
    [(mod (quot timestamp 60) 24) (mod timestamp 60)])
)

; This works as well
; (defn add-time [clock time]
;   (let [timestamp (abs (+ (apply timestamp clock) time))]
;     [(mod (quot timestamp 60) 24) (mod timestamp 60)])
; )

(defn add-time [clock time]
  (let [timestamp (abs (timestamp (first clock) (+ (second clock) time)))]
    [(mod (quot timestamp 60) 24) (mod timestamp 60)])
)