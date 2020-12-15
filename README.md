go_api

# Objetivos
 - [ ] CRUD de Cliente
    - Nome
    - CPF
    - EMAIL
    - Telefone
    - Endereço
 - [ ] CRUD de Produto
    - Nome
    - Preço
    - Peso
    - Imagem
    - Descrição
    - Estoque
    - Status
 - [ ] CRUD of Pedido
    - Codigo do Pedido
    - Cliente
    - Status
    - Subtotal dos Produtos
    - Valor do Frete
    - Total
    - Data de Criação
    - Forma de Pagamento
    - Forma de Entrega
 - [ ] CRUD do Item do Pedido
    - Codigo do Pedido
    - COdigo do Produto
    - Nome do Produto
    - Quantidade
    - Oreço Unitário
    - Total
 - [ ] Autenticação JWT
 - [ ] Apenas Produtos com Status ativo podem ser add ao Pedido
 - [ ] Apenas Produtos com Estoque disponivel podem ser add ao Pedido
 - [ ] Ao add um novo Produto ao Pedido o Total do Pedido deve ser atualizado
 - [ ] Apenas uma forma de entrega, o valor do frete é calculado da seguinte forma: Cada KG é add R$ 5,00, o valor é arredodado para cima.
 - [ ] Pagamento a vista tem 5% de desconto
 - [ ] Pagamento a prazo deve ser informado o numero de parcelas
 - [ ] Rota para acessar todos os pedidos realizados pelo cliente.
 - [ ] Teste Unitario para os Services do Sistema
 - [ ] Teste Automatizado para todas as Rotas do Sistema


 Obs: Deixar o Go executando: ``` CompileDaemon ```