/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40101 SET NAMES utf8 */;

use AprendaMais;

create table if not exists disciplina(
   id_disciplina integer not null auto_increment primary key,
   nome varchar(20) not null
)ENGINE =InnoDB DEFAULT CHARSET=latin1;