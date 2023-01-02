(ns rna-transcription)

(defn convert-to-rna [char]
  (case char
    \G "C"
    \C "G"
    \T "A"
    \A "U"
    (throw (AssertionError.))))

(defn to-rna [dna]
  (apply str (map #(convert-to-rna %) dna)))