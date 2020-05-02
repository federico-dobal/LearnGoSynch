Problem statement

There are 1 central variable that is updated regularly by a process and there are 2 more variables that needs to be updated by other processes. The goal is to have the 2 variables synchronised with the last update on central variable. We need to use Go routines iin order to syncronise the update.

The straigtforward solition would be to first execte the update and then trigger all the others individual updates. In sucha  case we do not need to use go routines. 
