package main

import "strings"

func QueryBuilder(guess Guess) {
	//Building the SQL Query based on the properties of the guessed champ

	//Gender
	if guess.Gender == 2 {
		sqlQuery += "AND `Gender`=(SELECT `Gender` FROM `Champs` WHERE `Name`='" + guess.Name + "') "
	} else {
		sqlQuery += "AND NOT (`Gender`=(SELECT `Gender` FROM `Champs` WHERE `Name`='" + guess.Name + "')) "
	}

	//Position
	temp := ""
	rows, err := db.Query("SELECT `Position` FROM `Champs` WHERE `Name`='" + guess.Name + "'")
	if err != nil {
		return
	}
	if rows.Next() {
		rows.Scan(&temp)
	}
	positions := strings.Split(temp, ",")

	if guess.Position == 2 {
		sqlQuery += "AND `Position`=(SELECT `Position` FROM `Champs` WHERE `Name`='" + guess.Name + "') "
	} else if guess.Position == 1 {
		sqlQuery += "AND "
		for i := 0; i < len(positions); i++ {
			sqlQuery += "`Position` LIKE '%" + positions[i] + "%' "
			if i != len(positions)-1 {
				sqlQuery += "OR "
			}
		}
	} else {
		sqlQuery += "AND NOT("
		for i := 0; i < len(positions); i++ {
			sqlQuery += "`Position` LIKE '%" + positions[i] + "%' "
			if i != len(positions)-1 {
				sqlQuery += "AND "
			}
		}
		sqlQuery += ") "
	}

	//Species
	rows, err = db.Query("SELECT `Species` FROM `Champs` WHERE `Name`='" + guess.Name + "'")
	if err != nil {
		return
	}
	if rows.Next() {
		rows.Scan(&temp)
	}
	species := strings.Split(temp, ",")

	if guess.Species == 2 {
		sqlQuery += "AND `Species`=(SELECT `Species` FROM `Champs` WHERE `Name`='" + guess.Name + "') "
	} else if guess.Species == 1 {
		sqlQuery += "AND "
		for i := 0; i < len(species); i++ {
			sqlQuery += "`Species` LIKE '%" + species[i] + "%' "
			if i != len(species)-1 {
				sqlQuery += "OR "
			}
		}
	} else {
		sqlQuery += "AND NOT("
		for i := 0; i < len(species); i++ {
			sqlQuery += "`Species` LIKE '%" + species[i] + "%' "
			if i != len(species)-1 {
				sqlQuery += "AND "
			}
		}
		sqlQuery += ") "
	}

	//Resource
	if guess.Resource == 2 {
		sqlQuery += "AND `Resource`=(SELECT `Resource` FROM `Champs` WHERE `Name`='" + guess.Name + "') "
	} else {
		sqlQuery += "AND NOT (`Resource`=(SELECT `Resource` FROM `Champs` WHERE `Name`='" + guess.Name + "')) "
	}

	//RangeType
	if guess.RangeType == 2 {
		sqlQuery += "AND `RangeType`=(SELECT `RangeType` FROM `Champs` WHERE `Name`='" + guess.Name + "') "
	} else {
		sqlQuery += "AND NOT (`RangeType`=(SELECT `RangeType` FROM `Champs` WHERE `Name`='" + guess.Name + "')) "
	}

	//Region
	rows, err = db.Query("SELECT `Region` FROM `Champs` WHERE `Name`='" + guess.Name + "'")
	if err != nil {
		return
	}
	if rows.Next() {
		rows.Scan(&temp)
	}
	region := strings.Split(temp, ",")

	if guess.Region == 2 {
		sqlQuery += "AND `Region`=(SELECT `Region` FROM `Champs` WHERE `Name`='" + guess.Name + "') "
	} else if guess.Region == 1 {
		sqlQuery += "AND "
		for i := 0; i < len(region); i++ {
			sqlQuery += "`Region` LIKE '%" + region[i] + "%' "
			if i != len(region)-1 {
				sqlQuery += "OR "
			}
		}
	} else {
		sqlQuery += "AND NOT("
		for i := 0; i < len(region); i++ {
			sqlQuery += "`Region` LIKE '%" + region[i] + "%' "
			if i != len(region)-1 {
				sqlQuery += "AND "
			}
		}
		sqlQuery += ") "
	}

	//ReleaseYear
	if guess.ReleaseYear == 2 {
		sqlQuery += "AND `ReleaseYear`=(SELECT `ReleaseYear` FROM `Champs` WHERE `Name`='" + guess.Name + "') "
	} else if guess.ReleaseYear == 1 {
		sqlQuery += "AND `ReleaseYear`>(SELECT `ReleaseYear` FROM `Champs` WHERE `Name`='" + guess.Name + "') "
	} else {
		sqlQuery += "AND `ReleaseYear`<(SELECT `ReleaseYear` FROM `Champs` WHERE `Name`='" + guess.Name + "') "
	}
}
