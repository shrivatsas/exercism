(ns armstrong-numbers)

(defn exp [n x]
  (reduce * (repeat n x)))

(defn armstrong? [num]
  (let [ncount (count (str num))]
  (->> num
       (str)
       (map #(Character/getNumericValue %))
       (map #(exp ncount %))
       (reduce +)
       (= num))))