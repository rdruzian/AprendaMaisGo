/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40101 SET NAMES utf8 */;

use AprendaMais;

create table if not exists questao(
    id_prova integer not null auto_increment primary key,
    id_faculdade integer not null,
    figura blob,
    ano datetime not null,
    constraint fk_id_faculdade foreign key (id_faculdade) references faculdade (id_faculdade)
)ENGINE =InnoDB DEFAULT CHARSET=latin1;