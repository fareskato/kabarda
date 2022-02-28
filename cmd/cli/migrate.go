package main

/*
	All needed logic for handling migrations
*/

// doMigrate run migrations commands
func doMigrate(arg2, arg3 string) error {
	// get the connection string
	dsn := getDSN()
	// run migration commands
	switch arg2 {
	case "up":
		err := kbr.MigrateUp(dsn)
		if err != nil {
			return err
		}
	case "down":
		// roll back all migration
		if arg3 == "all" {
			err := kbr.MigrateDownAll(dsn)
			if err != nil {
				return err
			}
			// roll back the last migration
		} else {
			err := kbr.MigrateSteps(-1, dsn)
			if err != nil {
				return err
			}
		}
	// reset DB: run all migrations down then run up again
	case "reset":
		err := kbr.MigrateDownAll(dsn)
		if err != nil {
			return err
		}
		err = kbr.MigrateUp(dsn)
		if err != nil {
			return err
		}
	default:
		showHelp()
	}
	return nil
}
