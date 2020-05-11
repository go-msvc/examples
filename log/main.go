package main

import (
	"os"

	"github.com/go-msvc/log"
)

func main() {

	//normal logging to terminal at debug level
	{
		log := log.New(os.Stderr)

		log.Tracef("Trace") //not written on debug level
		log.Debugf("Debug")
		log.Infof("Info")
		log.Warnf("Warn")
		log.Errorf("Error")
		log.Fatalf("Fatal")

		log.Trace("Trace") //not written on debug level
		log.Debug("Debug")
		log.Info("Info")
		log.Warn("Warn")
		log.Error("Error")
		log.Fatal("Fatal")

		//sub-logger:
		{
			ctxlog := log.New("ctx").With("context_id", "1234567890")
			ctxlog.Debug("Context Debug") //looks the same as above + logger name ".ctx"
		}
	}

	//write own log column format and add context_id to the log record
	{
		e := log.NewColumnEncoder().
			With(log.Column("time", log.TimeText("Jan 01"))).
			With(log.Column("myContextID", log.DataText("0x%08x", "context_id", 10))).
			With(log.Column("msg", log.MessageText(0)))
		log := log.New(os.Stderr).WithEncoder(e)
		log.Trace("Trace") //not written on debug level
		log.Debug("Debug")
		log.Info("Info")
		log.Set("context_id", 0x12345678)
		log.Debug("Debug")
		log.Info("Info")
	}

	//use default encoder + my context id
	{
		cols := log.DefaultEncoder().Columns()

		//insert before an existing column:
		idx := 0
		for ; idx < len(cols); idx++ {
			if cols[idx].Name() == "time" {
				break
			}
		}

		//add my column before level or append if not found
		newCols := []log.IColumn{}
		if idx < len(cols) {
			newCols = append(newCols, cols[0:idx]...)
			newCols = append(newCols, log.Column("myContextID", log.DataText("0x%08x", "context_id", 10)))
			newCols = append(newCols, cols[idx:]...)
		} else {
			newCols = append(newCols, cols...)
			newCols = append(newCols, log.Column("myContextID", log.DataText("0x%08x", "context_id", 10)))
		}
		e := log.NewColumnEncoder().With(newCols...)

		log := log.New(os.Stderr).WithEncoder(e)
		log.Trace("Trace") //not written on debug level
		log.Debug("Debug")
		log.Info("Info")
		log.Set("context_id", 0x12345678)
		log.Debug("Debug")
		log.Info("Info")
	}
}
