# Dados fictícios para apresentação

O ambiente de demonstração contém 40 voos, 51 passageiros, 255 bilhetes, 46 escalas, 14 aeronaves, 12 aeroportos, 12 cidades, 12 estados, três países e duas empresas. O gerente visualiza apenas os voos, a frota e os bilhetes da sua companhia. O atendente possui a visão de atendimento das duas empresas. Todos os registros adicionais são fictícios, os documentos usam o prefixo DM e os telefones ficam vazios.

A carga complementar roda automaticamente no início da aplicação somente com DEMO_MODE=true e com as três contas originais de demonstração presentes. A aplicação da carga ocorre em uma transação. Uma versão gravada no banco impede duplicação e preserva as alterações feitas durante a apresentação. Os registros anteriores são preservados. Se o armazenamento temporário do Render for recriado, a carga será criada novamente no início do serviço.

As datas dos novos voos programados cobrem os três dias seguintes à criação da carga, no fuso America/Sao_Paulo. Os voos ER 3101 a ER 3124 permitem demonstrar filtros, ordenação, paginação, tarifas diferentes, voos confirmados e planejados, emissão, check-in e cancelamento. O relógio do sistema atualiza automaticamente as situações conforme o horário real. Os exemplos de voo em andamento e embarque próximo evoluem com o tempo e não ficam congelados.

Entre com gerente@demo.local para apresentar planejamento de voos, frota, aeroportos, escalas, manutenção, consultas avançadas e auditoria. A senha dos três perfis de demonstração continua sendo Demo@2026!.

Os voos ER 3105 e ER 3113 possuem dois trechos, saindo de Guarulhos, passando por Brasília e chegando a Recife. O voo ER 3121 possui três trechos, passando também pelo Rio de Janeiro. Abra os detalhes de um deles para mostrar as escalas vinculadas, os bilhetes e o mapa de assentos.

Os voos ER 3201 a ER 3204 estão concluídos e possuem bilhetes usados. Os voos ER 3301 e ER 3302 foram cancelados e preservam os bilhetes cancelados e a justificativa na auditoria. O voo ER 3501 tem seus oito assentos ocupados para demonstrar indisponibilidade. Os bilhetes cancelados nos demais voos permitem mostrar que o assento fica disponível novamente.

A aeronave PRDM07 e o aeroporto FLN estão em manutenção com justificativa e histórico. A aeronave PRDM06 tem um registro de manutenção concluída. A PRDM08 está inativa e a PRDM09 está cadastrada, permitindo demonstrar diferentes fases do cadastro. O aeroporto GYN está cadastrado e pode ser ativado durante a apresentação. Há também uma escala planejada sem associação a voo.

Entre com atendente@demo.local para apresentar passageiros e bilhetes. Os passageiros Demo têm documentos sintéticos DM900000 a DM900047 e diferentes saldos de milhagem. Os quatro últimos estão cadastrados e ainda não possuem bilhetes. Escolha um voo futuro com assento disponível para emitir um bilhete. Abra um bilhete emitido, confirme o check-in e informe um portão como A01. O embarque só pode ser confirmado a partir de duas horas antes da partida.

Entre com passageiro@demo.local para apresentar a visão de Marina Oliveira. A conta possui bilhetes futuros adicionais e continua restrita ao seu próprio histórico. Use um bilhete emitido em voo futuro para demonstrar o check-in pelo passageiro. O voo EX 4101 pertence à outra empresa e permite verificar a separação da visão do gerente.

Nas consultas avançadas, a consulta de quantidade de registros resume a carga e as consultas de voos, escalas, passageiros e bilhetes apresentam resultados variados. Os registros vazios aparecem como Não informado. Os textos da interface e da carga complementar não usam travessões nem marcadores de lista.

Os testes automatizados verificam a integridade referencial, a capacidade dos assentos, a compatibilidade entre empresa e aeronave, a ausência de sobreposição de voos ativos por aeronave, a ordem das datas dos bilhetes, a ausência de duplicação, o bloqueio da carga fora do modo de demonstração e a reversão completa em caso de conflito.

Abra o site antes da apresentação para dar tempo de o serviço gratuito iniciar. O banco continua temporário no plano gratuito. A carga automática permite recuperar os cenários de demonstração após a recriação do armazenamento, mas alterações feitas manualmente podem ser perdidas nessa situação.
