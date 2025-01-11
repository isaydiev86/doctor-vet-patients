package db

const createPatientSQL = `
		INSERT INTO patient (fio, phone, animal, name, breed, gender, age, is_neutered)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`
const updatePatientSQL = `
    UPDATE patient
    SET fio = $1, phone = $2, animal = $3, name = $4, breed = $5, gender = $6, age = $7, is_neutered = $8
    WHERE id = $9;
`

const selectTreatmentDetailSQL = `
		SELECT 
            t.id, t.patient_id, t.doctor_id, t.temperature, t.status, t.created_at, t.updated_at, t.begin_at, t.end_at,
            t.comment, t.is_active, t.weight,
            p.id AS "patient.id", p.fio AS "patient.fio", p.phone AS "patient.phone",
            p.animal AS "patient.animal", p.name AS "patient.name", p.breed AS "patient.breed", p.gender AS "patient.gender",
            p.age AS "patient.age", p.is_neutered AS "patient.is_neutered",
            COALESCE(
                json_agg(
                    json_build_object(
                        'id', pr.id,
                        'treatment_id', pr.treatment_id,
                        'name', pr.name,
                        'course', pr.course,
                        'dose', pr.dose,
                        'category', pr.category,
                        'option', pr.option,
                        'created_at', pr.created_at,
                        'updated_at', pr.updated_at
                    )
                ) FILTER (WHERE pr.id IS NOT NULL), '[]'
            ) AS prescriptions
        FROM 
            treatment t
        LEFT JOIN 
            patient p ON t.patient_id = p.id
        LEFT JOIN 
            prescription pr ON t.id = pr.treatment_id
        WHERE 
            t.id = $1
        GROUP BY 
            t.id, p.id;
`
const selectReferenceSQL = `
	SELECT id, name, type
	FROM reference
	WHERE type = $1
`
const selectSymptomsSQL = `
	SELECT id, name
	FROM symptom;
`
const selectPreparationsSQL = `
	SELECT id, name, dose, course, category, option
	FROM preparation;
`
