Problem statement

Given a matrix build  an ew one where each element is the average of all its neigbours. Use Go routines and any synch utils proviided by Go.

It means new matrix element at posiition (i, j) will be:

(m(i-1, j-1) + m(i-1, j) + m(i-1, j+1) +
m(i, j-1) + m(i, j+1) + 
m(i+1, j-1) + m(i+1, j) + m(i+1, j+1)) / 8

It is required to consider out of bounds elements and remove from the summe but still calculate the average properly.

