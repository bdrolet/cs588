# UIC CS 588
Sergio Barajas
Ben Drolet

Shamir's Secret Sharing

Our implementation of Shamirs secret uses UINT64 encryption. We create 8 byte shares based on the secret passed in input. This is much more secure than lower bit value encoding implementations. We first set up the file structure. We have three files main.go, shareManager.go and algorithm.go. In main.go we take in the arguments for our program and do some simple error checking. We also call our code to createShares and write out to a file. In algorithm.go we have our code to calculate the shares and create them based on shamir's secret algorithm. In shareManager.go we have the code wich takes in the user defined arguments k, N and the secret and creates the shares to pass to the code in algorithm.go to calculate and create the shares. the code to later combine the shares will also be in shareManager.go. To run the code you need to build the go environment and then specify the K, N and secret values as arguments in that order.
