const API_BASE_URL = 'http://localhost:8080';
const searchInput = document.getElementById('search-cpf-input'); // Pegamos o input no index.html e registramos aqui
const searchButton = document.getElementById('search-btn');
const deleteButton = document.getElementById('delete-btn');
const resultDisplay = document.getElementById('result-display');
const userForm = document.getElementById('user-form');
const submitButton = document.getElementById('submit-btn');
const clearButton = document.getElementById('clear-btn');
const loadAllButton = document.getElementById('load-all-btn');
const usersTableBody = document.getElementById('users-table-body');
const nomeInput = document.getElementById('nome');
const cpfCnpjInput = document.getElementById('cpf_cnpj');
const emailInput = document.getElementById('email');
const agenciaInput = document.getElementById('agencia');
const contaInput = document.getElementById('conta');
const bancoInput = document.getElementById('banco');
const pixInput = document.getElementById('pix');
const findUserByCpf = async () => {
    const cpf = searchInput.value;
    if (!cpf) {
        resultDisplay.textContent = 'Por favor, digite um CPF.';
        return;
    }
    resultDisplay.textContent = 'Buscando...';
    try {
        const response = await fetch(`${API_BASE_URL}/users/${cpf}`); // Requisita p api o /users/cpf pra fazer um get
        if (!response.ok) {
            const errorResult = await response.json();
            throw new Error(errorResult.Error || errorResult.error || 'Usuário não encontrado.');
        }
        const user = await response.json();
        resultDisplay.textContent = JSON.stringify(user, null, 2);
        fillForm(user);
    }
    catch (error) {
        resultDisplay.textContent = `Falha na busca: ${error}`;
    }
};
const deleteUserByCpf = async () => {
    const cpf = searchInput.value;
    if (!cpf) {
        resultDisplay.textContent = 'Por favor, digite um CPF para deletar.';
        return;
    }
    if (!confirm(`Tem certeza que deseja deletar o usuário com CPF ${cpf}?`))
        return;
    resultDisplay.textContent = `Deletando usuário com CPF ${cpf}...`;
    try {
        const response = await fetch(`${API_BASE_URL}/users/${cpf}`, { method: 'DELETE' });
        const result = await response.json();
        if (!response.ok) {
            throw new Error(result.Error || result.error || 'Erro ao deletar usuário');
        }
        resultDisplay.textContent = result.Message || result.message || 'Usuário deletado com sucesso!';
        fetchAllUsers();
    }
    catch (error) {
        resultDisplay.textContent = `Falha ao deletar: ${error}`;
    }
};
const fetchAllUsers = async () => {
    usersTableBody.innerHTML = '<tr><td colspan="4">Carregando...</td></tr>';
    try {
        const response = await fetch(`${API_BASE_URL}/users`);
        if (!response.ok)
            throw new Error('Erro ao carregar usuários.');
        const users = await response.json();
        usersTableBody.innerHTML = '';
        if (users && users.length > 0) {
            users.forEach(user => {
                const row = usersTableBody.insertRow();
                row.innerHTML = `
                    <td>${user.id}</td>
                    <td>${user.name}</td>
                    <td>${user.cpf_cnpj}</td>
                    <td>${user.email}</td>
                    <td>${user.conta_bancaria.agencia}</td>
                    <td>${user.conta_bancaria.conta}</td>
                    <td>${user.conta_bancaria.banco}</td>
                    <td>${user.conta_bancaria.pix}</td>
                    <td>${new Date(user.created_at).toLocaleString()}</td>
                `;
            });
        }
        else {
            usersTableBody.innerHTML = '<tr><td colspan="4">Nenhum usuário encontrado.</td></tr>';
        }
    }
    catch (error) {
        usersTableBody.innerHTML = `<tr><td colspan="4">${error}</td></tr>`;
    }
};
const handleFormSubmit = async (event) => {
    event.preventDefault();
    if (cpfCnpjInput.disabled) {
        handleUpdateUser();
    }
    else {
        handleCreateUser();
    }
};
const handleCreateUser = async () => {
    const userData = {
        name: nomeInput.value,
        cpf_cnpj: cpfCnpjInput.value,
        email: emailInput.value,
        conta_bancaria: {
            agencia: agenciaInput.value,
            conta: contaInput.value,
            banco: bancoInput.value,
            pix: pixInput.value,
        },
    };
    resultDisplay.textContent = 'Criando usuário...';
    try {
        const response = await fetch(`${API_BASE_URL}/users`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(userData),
        });
        const result = await response.json();
        if (!response.ok) {
            throw new Error(result.Error || result.error || 'Erro ao criar usuário');
        }
        resultDisplay.textContent = result.Message || result.message || 'Usuário criado com sucesso!';
        clearForm();
        fetchAllUsers();
    }
    catch (error) {
        resultDisplay.textContent = `Falha ao criar usuário: ${error}`;
    }
};
const handleUpdateUser = async () => {
    const cpf = cpfCnpjInput.value;
    const updateData = {
        name: nomeInput.value,
        email: emailInput.value,
        conta_bancaria: {
            agencia: agenciaInput.value,
            conta: contaInput.value,
            banco: bancoInput.value,
            pix: pixInput.value,
        },
    };
    resultDisplay.textContent = 'Atualizando usuário...';
    try {
        const response = await fetch(`${API_BASE_URL}/users/${cpf}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(updateData),
        });
        const result = await response.json();
        if (!response.ok) {
            throw new Error(result.Error || result.error || 'Erro ao atualizar usuário');
        }
        resultDisplay.textContent = result.Message || result.message || 'Usuário atualizado com sucesso!';
        clearForm();
        fetchAllUsers();
    }
    catch (error) {
        resultDisplay.textContent = `Falha ao atualizar usuário: ${error}`;
    }
};
const fillForm = (user) => {
    nomeInput.value = user.name;
    cpfCnpjInput.value = user.cpf_cnpj;
    emailInput.value = user.email;
    agenciaInput.value = user.conta_bancaria.agencia;
    contaInput.value = user.conta_bancaria.conta;
    bancoInput.value = user.conta_bancaria.banco;
    pixInput.value = user.conta_bancaria.pix;
    submitButton.textContent = 'Atualizar Usuário';
    cpfCnpjInput.disabled = true;
};
const clearForm = () => {
    userForm.reset();
    submitButton.textContent = 'Criar Usuário';
    cpfCnpjInput.disabled = false;
    nomeInput.disabled = false;
    resultDisplay.textContent = 'Aguardando ação...';
};
searchButton.addEventListener('click', findUserByCpf);
deleteButton.addEventListener('click', deleteUserByCpf);
loadAllButton.addEventListener('click', fetchAllUsers);
userForm.addEventListener('submit', handleFormSubmit);
clearButton.addEventListener('click', clearForm);
document.addEventListener('DOMContentLoaded', fetchAllUsers);
export {};
//# sourceMappingURL=index.js.map