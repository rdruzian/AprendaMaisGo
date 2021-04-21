/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40101 SET NAMES utf8 */;

use AprendaMais;

create table if not exists aluno(
    id_aluno integer not null auto_increment primary key,
    nome varchar(100) not null,
    sexo varchar(1) not null ,
    uf varchar(2) not null,
    id_area integer not null,
    constraint fk_id_area foreign key (id_area) references area (id_area)
)ENGINE =InnoDB DEFAULT CHARSET=latin1;