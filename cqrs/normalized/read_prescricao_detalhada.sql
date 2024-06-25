drop table if exists vw_prescricoes_medicamentos_detalhadas;

CREATE TABLE IF NOT EXISTS vw_prescricoes_medicamentos_detalhadas (
	id SERIAL primary key,
    id_prescricao INT,
    data_prescricao TIMESTAMP NOT NULL,
    id_medico INT NOT NULL,
    nome_medico VARCHAR(255) NOT NULL,
    especialidade_medico VARCHAR(255) NOT NULL,
    id_paciente INT NOT NULL,
    nome_paciente VARCHAR(255) NOT NULL,
    data_nascimento_paciente DATE NOT NULL,
    endereco_paciente VARCHAR(255),
    id_medicamento INT NOT NULL,
    nome_medicamento VARCHAR(255) NOT NULL,
    descricao_medicamento TEXT,
    horario VARCHAR(50) NOT NULL,
    dosagem VARCHAR(50) NOT null,
    FOREIGN KEY (id_medico) REFERENCES Medicos(id),
    FOREIGN KEY (id_paciente) REFERENCES Pacientes(id),
    FOREIGN KEY (id_medicamento) REFERENCES Medicamentos(id),
    FOREIGN KEY (id_prescricao) REFERENCES Prescricoes(id)
);

INSERT INTO vw_prescricoes_medicamentos_detalhadas (
    id_prescricao,
    data_prescricao,
    id_medico,
    nome_medico,
    especialidade_medico,
    id_paciente,
    nome_paciente,
    data_nascimento_paciente,
    endereco_paciente,
    id_medicamento,
    nome_medicamento,
    descricao_medicamento,
    horario,
    dosagem
)
SELECT
    p.id AS id_prescricao,
    p.data_prescricao,
    m.id AS id_medico,
    m.nome AS nome_medico,
    m.especialidade AS especialidade_medico,
    pac.id AS id_paciente,
    pac.nome AS nome_paciente,
    pac.data_nascimento AS data_nascimento_paciente,
    pac.endereco AS endereco_paciente,
    med.id AS id_medicamento,
    med.nome AS nome_medicamento,
    med.descricao AS descricao_medicamento,
    pm.horario,
    pm.dosagem
FROM
    Prescricoes p
    JOIN Medicos m ON p.id_medico = m.id
    JOIN Pacientes pac ON p.id_paciente = pac.id
    JOIN Prescricao_Medicamentos pm ON p.id = pm.id_prescricao
    JOIN Medicamentos med ON pm.id_medicamento = med.id;