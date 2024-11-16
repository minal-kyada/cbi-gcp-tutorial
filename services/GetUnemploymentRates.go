package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/models"
	"net/http"
)

func GetUnemploymentRates(db *sql.DB) {
	fmt.Println("GetUnemploymentRates: Collecting Unemployment Rates Data")

	// This function is NOT complete
	// It provides code-snippets for the data source: https://data.cityofchicago.org/Health-Human-Services/Public-Health-Statistics-Selected-public-health-in/iqnk-2tcu/data

	// Data Collection needed from two data sources:
	// 1. https://data.cityofchicago.org/Health-Human-Services/Public-Health-Statistics-Selected-public-health-in/iqnk-2tcu/data

	drop_table := `drop table if exists unemployment`
	_, err := db.Exec(drop_table)
	if err != nil {
		panic(err)
	}

	create_table := `CREATE TABLE IF NOT EXISTS "unemployment" (
						"id"   SERIAL , 
						"community_area" VARCHAR(255) UNIQUE, 
						"community_area_name" VARCHAR(255), 
						"birth_rate" VARCHAR(255), 
						"general_fertility_rate" VARCHAR(255), 
						"low_birth_weight" VARCHAR(255),
						
						"prenatal_care_beginning_in_first_trimester" VARCHAR(255) , 
						"preterm_births" VARCHAR(255), 
						"teen_birth_rate" VARCHAR(255), 
						"assault_homicide" VARCHAR(255), 
						"breast_cancer_in_females" VARCHAR(255),
						
						"cancer_all_sites" VARCHAR(255) , 
						"colorectal_cancer" VARCHAR(255), 
						"diabetes_related" VARCHAR(255), 
						"firearm_related" VARCHAR(255), 
						"infant_mortality_rate" VARCHAR(255),
						
						"lung_cancer" VARCHAR(255) , 
						"prostate_cancer_in_males" VARCHAR(255), 
						"stroke_cerebrovascular_disease" VARCHAR(255), 
						"childhood_blood_lead_level_screening" VARCHAR(255), 
						"childhood_lead_poisoning" VARCHAR(255),
						
						"gonorrhea_in_females" VARCHAR(255) , 
						"gonorrhea_in_males" VARCHAR(255), 
						"tuberculosis" VARCHAR(255), 
						"below_poverty_level" VARCHAR(255), 
						"crowded_housing" VARCHAR(255),
						
						"dependency" VARCHAR(255) , 
						"no_high_school_diploma" VARCHAR(255), 
						"unemployment" VARCHAR(255), 
						"per_capita_income" VARCHAR(255),
						PRIMARY KEY ("id") 
					);`

	_, _err := db.Exec(create_table)
	if _err != nil {
		panic(_err)
	}

	fmt.Println("Created Table for Unemployment")

	// While doing unit-testing keep the limit value to 500
	// later you could change it to 1000, 2000, 10,000, etc.
	var url = "https://data.cityofchicago.org/resource/iqnk-2tcu.json?$limit=100"

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Received data from SODA REST API for Unemployment")

	body, _ := ioutil.ReadAll(res.Body)
	var unemployment_data_list models.UnemploymentJsonRecords
	json.Unmarshal(body, &unemployment_data_list)

	// prettyJSON, _ := json.MarshalIndent(unemployment_data_list, "", "    ")
	// fmt.Println("Building Permit Data: ", string(prettyJSON))

	for i := 0; i < len(unemployment_data_list); i++ {

		// We will execute defensive coding to check for messy/dirty/missing data values
		// Any record that has messy/dirty/missing data we don't enter it in the data lake/table

		community_area := unemployment_data_list[i].Community_area
		if community_area == "" {
			continue
		}

		community_area_name := unemployment_data_list[i].Community_area_name
		if community_area_name == "" {
			continue
		}

		birth_rate := unemployment_data_list[i].Birth_rate
		if birth_rate == "" {
			continue
		}

		general_fertility_rate := unemployment_data_list[i].General_fertility_rate
		if general_fertility_rate == "" {
			continue
		}

		low_birth_weight := unemployment_data_list[i].Low_birth_weight
		if low_birth_weight == "" {
			continue
		}

		prenatal_care_beginning_in_first_trimester := unemployment_data_list[i].Prenatal_care_beginning_in_first_trimester
		if prenatal_care_beginning_in_first_trimester == "" {
			continue
		}

		preterm_births := unemployment_data_list[i].Preterm_births
		if preterm_births == "" {
			continue
		}

		teen_birth_rate := unemployment_data_list[i].Teen_birth_rate
		if teen_birth_rate == "" {
			continue
		}

		assault_homicide := unemployment_data_list[i].Assault_homicide
		if assault_homicide == "" {
			continue
		}

		breast_cancer_in_females := unemployment_data_list[i].Breast_cancer_in_females
		if breast_cancer_in_females == "" {
			continue
		}

		cancer_all_sites := unemployment_data_list[i].Cancer_all_sites
		if cancer_all_sites == "" {
			continue
		}

		colorectal_cancer := unemployment_data_list[i].Colorectal_cancer
		if colorectal_cancer == "" {
			continue
		}

		diabetes_related := unemployment_data_list[i].Diabetes_related
		if diabetes_related == "" {
			continue
		}

		firearm_related := unemployment_data_list[i].Firearm_related
		if firearm_related == "" {
			continue
		}

		infant_mortality_rate := unemployment_data_list[i].Infant_mortality_rate
		if infant_mortality_rate == "" {
			continue
		}

		lung_cancer := unemployment_data_list[i].Lung_cancer
		if lung_cancer == "" {
			continue
		}

		prostate_cancer_in_males := unemployment_data_list[i].Prostate_cancer_in_males
		if prostate_cancer_in_males == "" {
			continue
		}

		stroke_cerebrovascular_disease := unemployment_data_list[i].Stroke_cerebrovascular_disease
		if stroke_cerebrovascular_disease == "" {
			continue
		}

		childhood_blood_lead_level_screening := unemployment_data_list[i].Childhood_blood_lead_level_screening
		if childhood_blood_lead_level_screening == "" {
			continue
		}

		childhood_lead_poisoning := unemployment_data_list[i].Childhood_lead_poisoning
		if childhood_lead_poisoning == "" {
			continue
		}

		gonorrhea_in_females := unemployment_data_list[i].Gonorrhea_in_females
		if gonorrhea_in_females == "" {
			continue
		}

		gonorrhea_in_males := unemployment_data_list[i].Gonorrhea_in_males
		if gonorrhea_in_males == "" {
			continue
		}

		tuberculosis := unemployment_data_list[i].Tuberculosis
		if tuberculosis == "" {
			continue
		}

		below_poverty_level := unemployment_data_list[i].Below_poverty_level
		if below_poverty_level == "" {
			continue
		}

		crowded_housing := unemployment_data_list[i].Crowded_housing
		if crowded_housing == "" {
			continue
		}

		dependency := unemployment_data_list[i].Dependency
		if dependency == "" {
			continue
		}

		no_high_school_diploma := unemployment_data_list[i].No_high_school_diploma
		if no_high_school_diploma == "" {
			continue
		}

		per_capita_income := unemployment_data_list[i].Per_capita_income
		if per_capita_income == "" {
			continue
		}

		unemployment := unemployment_data_list[i].Unemployment
		if unemployment == "" {
			continue
		}

		sql := `INSERT INTO unemployment ("community_area" , 
		"community_area_name" , 
		"birth_rate" , 
		"general_fertility_rate" , 
		"low_birth_weight" ,
		
		"prenatal_care_beginning_in_first_trimester" , 
		"preterm_births" , 
		"teen_birth_rate" , 
		"assault_homicide" , 
		"breast_cancer_in_females" ,
		
		"cancer_all_sites"  , 
		"colorectal_cancer" , 
		"diabetes_related" , 
		"firearm_related" , 
		"infant_mortality_rate" ,
		
		"lung_cancer" , 
		"prostate_cancer_in_males" , 
		"stroke_cerebrovascular_disease" , 
		"childhood_blood_lead_level_screening" , 
		"childhood_lead_poisoning" ,
		
		"gonorrhea_in_females"  , 
		"gonorrhea_in_males" , 
		"tuberculosis" , 
		"below_poverty_level" , 
		"crowded_housing" ,
		
		"dependency"  , 
		"no_high_school_diploma" , 
		"unemployment" , 
		"per_capita_income" )
		values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,$11, $12, $13, $14, $15,$16, $17, $18, $19, $20,$21, $22, $23, $24, $25,$26, $27, $28, $29)`

		_, err = db.Exec(
			sql,
			community_area,
			community_area_name,
			birth_rate,
			general_fertility_rate,
			low_birth_weight,

			prenatal_care_beginning_in_first_trimester,
			preterm_births,
			teen_birth_rate,
			assault_homicide,
			breast_cancer_in_females,

			cancer_all_sites,
			colorectal_cancer,
			diabetes_related,
			firearm_related,
			infant_mortality_rate,

			lung_cancer,
			prostate_cancer_in_males,
			stroke_cerebrovascular_disease,
			childhood_blood_lead_level_screening,
			childhood_lead_poisoning,

			gonorrhea_in_females,
			gonorrhea_in_males,
			tuberculosis,
			below_poverty_level,
			crowded_housing,

			dependency,
			no_high_school_diploma,
			unemployment,
			per_capita_income)

		if err != nil {
			panic(err)
		}

	}

}
