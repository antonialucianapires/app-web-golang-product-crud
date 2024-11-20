-- Criação da tabela
CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    descricao VARCHAR(255) NOT NULL,
    preco DECIMAL NOT NULL,
    quantidade INTEGER NOT NULL,
    data_criacao TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Inserção dos produtos iniciais
INSERT INTO produtos (nome, descricao, preco, quantidade, data_criacao) VALUES
('Camiseta', 'Camiseta de algodão, ideal para o dia a dia', 49.90, 12, NOW()),
('Calça Jeans', 'Calça jeans confortável com corte reto', 89.90, 8, NOW()),
('Tênis Casual', 'Tênis casual, ideal para passeios e trabalho', 129.90, 5, NOW());

