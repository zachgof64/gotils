package gotils

//Check will log output to console on error
func Check(err error) {
	if(err != nil){
		panic(err)
	}
} 