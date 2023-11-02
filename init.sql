CREATE DATABASE IF NOT EXISTS APIExercise;

USE APIExercise;

CREATE TABLE IF NOT EXISTS alerts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    alert_id VARCHAR(32),
    service_id VARCHAR(32),
    service_name VARCHAR(32),
    model VARCHAR(32),
    alert_type VARCHAR(32),
    alert_ts VARCHAR(32),
    alert_ts_int bigint,
    severity VARCHAR(32),
    team_slack VARCHAR(32)
);

/*
insert into alerts (alert_id, service_id, service_name, model, alert_type, alert_ts, severity, team_slack)
select 'My alert_id', 'My service_id', 'My service_name', 'My model', 'My alert_type', '2023-10-31 15:30:00', 'My severity', 'My team_slack"';
*/
