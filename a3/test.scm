
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
(count-primes 10000)





