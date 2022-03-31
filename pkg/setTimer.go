package pkg

import "flag"

// Init Document
/*
 *--------------------------------------------
 * @brief    Init
 *           setTimerInitialization
 *
 * @param[in] int time
 *
 * @return
 *
 * @rules
 *--------------------------------------------
 */
func Init(t int) {
	flag.IntVar(&t, "time", 0, "timeLimit")
	flag.Parse()
}
