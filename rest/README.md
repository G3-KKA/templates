



__To add new method__

# package domain
1) Create single-method interface
2) Embed the interface into Repository interface
# package handlers
3) Create method-named 'State 
    - Get logic as repo's method
    - Incapsulates it via single-method interface
    - Implements http.Handler
    ,  whithin which the logic is invoked
4) Create constructor for 'State
    - Get needed part of logic as repo
# package clientrepo
5) Finally, implement actual logic  

