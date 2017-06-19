
(load "a3.scm")

(my-last '(cat))
(my-last '(cat dog))
(my-last '(cat dog (1 2 3)))
(my-last '())

(snoc 'a '())
(snoc 'a '(1))
(snoc 'a '(1 2))
(snoc 'a '(1 2 3))
(snoc '(1 2 3) '(1 2 3))

(range 0)
(range -4)
(range 4)
(range 9)

(deep-sum '())
(deep-sum '(a))
(deep-sum '(2))
(deep-sum '(a 2))
(deep-sum '(a 2 3))
(deep-sum '(a 2 (b 2) 3))
(deep-sum '(a 2 (b (1 c)) 3))

(count-primes -10)
(count-primes 0)
(count-primes 5)
(count-primes 10)
(count-primes 100)
(count-primes 1000)
;(count-primes 10000)

(is-bit? 0)
(is-bit? 1)
(is-bit? 2)
(is-bit? '(0 1))
(is-bit? 'cow)

(is-bit-seq? '())
(is-bit-seq? '(0))
(is-bit-seq? '(1))
(is-bit-seq? '(2))
(is-bit-seq? '(0 1))
(is-bit-seq? '(0 2))
(is-bit-seq? '(0 1 1 1 1 1 0 1 1 0 1 0))
(is-bit-seq? '(0 1 1 1 1 0 0 1 1 0 1 5))
(is-bit-seq? '((0 1) 1 1 1 1 0 0 1 1 0 1 0))
(is-bit-seq? '(cow 1 1 1 1 0 0 1 1 0 1 0))
(is-bit-seq? '(1 1 1 1 1 0 1 1 0 1 (dog cat)))

(all-bit-seqs 0)
(all-bit-seqs 1)
(all-bit-seqs 2)


