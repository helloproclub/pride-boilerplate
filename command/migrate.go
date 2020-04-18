package command

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/helloproclub/pride-boilerplate/helper"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
	Long:  `This subcommand used to migrate database`,
	Run:   processMigration,
}

func init() {
	migrateCmd.PersistentFlags().Int("step", 0, "maximum migration steps")
	migrateCmd.PersistentFlags().String("direction", "up", "migration direction")
	RootCmd.AddCommand(migrateCmd)
}

func processMigration(cmd *cobra.Command, args []string) {

	direction := cmd.Flag("direction").Value.String()
	stepStr := cmd.Flag("step").Value.String()
	step, _ := strconv.Atoi(stepStr)

	migrations := &migrate.FileMigrationSource{
		Dir: "./migration",
	}

	migrate.SetTable("schema_migrations")
	dbURL := helper.GetEnv("DB_URL", "").(string)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(dbURL + " failed to connect")
		panic(err.Error())
	}

	var n int
	if direction == "down" {
		n, err = migrate.ExecMax(db, "postgres", migrations, migrate.Down, step)
	} else {
		n, err = migrate.ExecMax(db, "postgres", migrations, migrate.Up, step)
	}
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Applied %d migrations!\n", n)

}
