# aRPC (Antenna RPC)

[![project_badge](https://img.shields.io/badge/CodeFairyUFRJ/Universe-black.svg?style=for-the-badge&logo=github "Project Badge")](https://github.com/CodeFairyUFRJ/Universe)
[![license_badge](https://img.shields.io/github/license/CodeFairyUFRJ/Universe.svg?style=for-the-badge& "License Badge")](http://creativecommons.org/licenses/by-sa/4.0/)

Um framework de chamada de procedimento remoto (RPC) para uso em computação de alto desempenho (HPC)

## Resumo
A crescente adoção da arquitetura de microsserviços e a necessidade de comunicação entre linguagens distintas estimulou o desenvolvimento de novas soluções para a chamada de procedimento remoto (RPC). A diversidade de necessidades e propósitos resultou em uma variedade de implementações de RPC: algumas com foco na ergonomia de software; outras na abrangência de linguagens e funcionalidades; e, finalmente, uma parcela visando a eficiência em computação de alto desempenho (HPC). Nesse sentido, é apresentado neste trabalho um framework de RPC, de nome aRPC, com ênfase tanto no desempenho como na ergonomia de software, inspirado no framework gRPC, e que faz uso de novos serializadores e do protocolo de transporte QUIC para comunicação. 
Nas avaliações efetuadas, o aRPC obteve desempenho superior ao gRPC nos casos com grande quantidades de elementos nas estruturas de dados e quando os dados são mais heterogêneos e menos sintéticos. O protocolo proposto é capaz de oferecer desempenho até 7% superior ao gRPC, desde que as premissas descritas sejam respeitadas. Em situações com perda frequente de pacotes ou em redes de baixa qualidade, o aRPC possui desempenho muito superior ao gRPC, sendo até três vezes melhor em alguns testes. O desempenho do aRPC abre um campo de aplicação em sistemas de computação alto desempenho e a resiliência apresentada faz com que seja uma opção interessante nos ambientes de IoT. Entretanto, o protocolo gRPC apresenta melhor desempenho para algumas estruturas de dados específicas e para volumes de dados reduzidos. 

## Integrantes
+ Ericson Soares ([@fogodev](https://github.com/fogodev))
+ Raphael Almeida ([@Raphael-C-Almeida](https://github.com/Raphael-C-Almeida))
+ Vítor Vasconcellos ([@HeavenVolkoff](https://github.com/HeavenVolkoff))

## Licensa
[![Logo cc-by-sa-4.0](https://i.creativecommons.org/l/by-sa/4.0/88x31.png)](http://creativecommons.org/licenses/by-sa/4.0/)

Este trabalho está licenciado sob uma [Licensa Creative Commons Attribution-ShareAlike 4.0 International](http://creativecommons.org/licenses/by-sa/4.0/).
