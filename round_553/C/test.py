import math, sys
import random
import subprocess

N = 20000

n = [i for i in xrange(1, N*2)]
n1 = filter(lambda x: x%2 == 0, n)
n2 = filter(lambda x: x%2 == 1, n)

s1, s2 = [], []
s = []


c1, c2 = 0, 0
i = 1

while i <= N:
	if int(math.log(i, 2)) % 2 == 0:
		popped, n2 = n2[0], n2[1:]
		s1.append(popped)
		s.append(popped)
	else:
		popped, n1 = n1[0], n1[1:]
		s2.append(popped)
		s.append(popped)
	i += 1

f = open('testcases', 'wb')

for c in xrange(5000):
	l = random.randint(1, N)
	r = random.randint(l, N)
	res = sum(s[l-1:r])
	go_res = int(subprocess.check_output(['go', 'run', 'solver2.go', str(l), str(r)]))

	if c % 100 == 0:
		print "[+]finished {} tests".format(c)
	if res != go_res:
		print "l, r:", l, r
		print "res:", res
		print "go res:", go_res
