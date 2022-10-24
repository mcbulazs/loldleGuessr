CREATE DATABASE `Loldle` DEFAULT CHARACTER SET UTF8MB4 COLLATE UTF8MB4_HUNGARIAN_CI;

CREATE TABLE `Champs` (
	`Name` VARCHAR(255) NOT NULL,
	`Gender` ENUM('Male', 'Female', 'Other') NOT NULL,
	`Position` SET('Top', 'Jungle', 'Mid', 'Bottom', 'Support') NOT NULL,
	`Species` SET('Unknown', 'Aspect', 'Human', 'Magicborn', 'Yordle', 'Magically Altered', 'Golem', 'Cat', 'God-Warrior', 'Vastayan', 'Chemically Altered', 'Cyborg', 'God', 'Spirit', 'Undead', 'Void-Being', 'Darkin', 'Spiritualist', 'Iceborn', 'Rat', 'Troll', 'Celestial', 'Brackern', 'Revenant', 'Dragon', 'Demon', 'Minotaur') NOT NULL,
	`Resource` ENUM('Mana', 'Energy', 'Manaless', 'Fury', 'Rage', 'Shield', 'Heat', 'Health costs', 'Bloodthirst', 'Courage', 'Ferocity', 'Grit', 'Flow') NOT NULL,
	`Range Type` ENUM('Melee', 'Ranged', 'Both') NOT NULL,
	`Region` SET('Ixtal', 'Targon', 'Icathia', 'Runeterra', 'Shurima', 'Piltover', 'Zaun', 'Ionia', 'Bandle City', 'Blessed Isles', 'Shadow Isles', 'Demacia', 'Noxus', 'Freljord', 'Camavor', 'Void', 'Bilgewater') NOT NULL,
	`Release Year` INT(4) NOT NULL,
	PRIMARY KEY(`Name`)	
);