package db

const createPatientSQL = `
		INSERT INTO patient (fio, phone, address, animal, name, breed, gender, age, is_neutered)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`
const updatePatientSQL = `
    UPDATE patient
    SET fio = $1, phone = $2, address = $3, animal = $4, name = $5, breed = $6, gender = $7, age = $8, is_neutered = $9
    WHERE id = $10;
`

var selectTreatmentsSQL = `
		SELECT 
            t.id, t.patient_id, t.doctor_id, t.temperature, t.status, t.created_at, t.updated_at, t.begin_at, t.end_at, 
            t.comment, t.is_active, t.weight,
            p.id AS "patient.id", p.fio AS "patient.fio", p.phone AS "patient.phone",
            p.age AS "patient.age", p.address AS "patient.address", p.animal AS "patient.animal", p.name AS "patient.name", 
            p.breed AS "patient.breed", p.gender AS "patient.gender", p.is_neutered AS "patient.is_neutered"
        FROM 
            treatment t
        LEFT JOIN 
            patient p ON t.patient_id = p.id
        WHERE 1=1`

const selectTreatmentDetailSQL = `
		SELECT 
            t.id, t.patient_id, t.doctor_id, t.temperature, t.status, t.created_at, t.updated_at, t.begin_at, t.end_at,
            t.comment, t.is_active, t.weight,
            p.id AS "patient.id", p.fio AS "patient.fio", p.phone AS "patient.phone", p.address AS "patient.address",
            p.animal AS "patient.animal", p.name AS "patient.name", p.breed AS "patient.breed", p.gender AS "patient.gender",
            p.age AS "patient.age", p.is_neutered AS "patient.is_neutered",
            COALESCE(
                json_agg(
                    json_build_object(
                        'id', pr.id,
                        'treatment_id', pr.treatment_id,
                        'preparation', pr.preparation,
                        'course', pr.course,
                        'dose', pr.dose,
                        'amount', pr.amount,
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
