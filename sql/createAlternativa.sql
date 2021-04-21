/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40101 SET NAMES utf8 */;

use AprendaMais;

create table if not exists questao(
    id_questao integer not null auto_increment primary key,
    id_disciplina integer not null,
    texto_alternativa varchar(65535) not null,
    texto varchar(65535),
    a blob not null,
    b blob not null,
    c blob not null,
    d blob not null,
    e blob,
    resp_correta char(1) not null,
    CONSTRAINT fk_id_disciplina FOREIGN KEY (id_disciplina) REFERENCES disciplina (id_disciplina)
)ENGINE =InnoDB DEFAULT CHARSET=latin1;