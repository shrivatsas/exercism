(ns beer-song)
    (use '[clojure.string :only (join)])

(defn get-beer-bottles [count]
    (case count
        0 "no more bottles"
        1 "1 bottle"
        (format "%d bottles" count)))

(defn get-beer-on-wall [count]
    (str (get-beer-bottles count) " of beer on the wall"))

(defn get-beer [count]
    (case count
        0 "Go to the store and buy some more"
        1 "Take it down and pass it around"
        "Take one down and pass it around"))

(defn verse-number [init, left]
    (str (clojure.string/capitalize (get-beer-on-wall init)) ", " (get-beer-bottles init) " of beer.\n"
       (get-beer init) ", " (get-beer-on-wall left) ".\n"))
       

(defn verse [count]
    (case count
        0 (verse-number count 99)
        (verse-number count (dec count))))

(defn sing 
    ([start] (sing start 0))
    ([start end]
        (join "\n"
            (map #(verse %) (range start (dec end) -1)))))