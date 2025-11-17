const modal = document.getElementById('editModal');
if(modal){
    modal.addEventListener('show.bs.modal', event =>{

        const button = event.relatedTarget;

        const id = button.getAttribute('data-bs-id');
        const name = button.getAttribute('data-bs-nome');
        const email = button.getAttribute('data-bs-email');
        const telefone = button.getAttribute('data-bs-telefone');

        const inputId = document.getElementById('editId');
        const inputName = document.getElementById('editName');
        const inputEmail = document.getElementById('editEmail');
        const inputTelefone = document.getElementById('editPhone');

        inputId.value = id;
        inputName.value = name;
        inputEmail.value = email;
        inputTelefone.value = telefone;
    });
}

const deleteModal = document.getElementById('deleteModal');
if(deleteModal){
    deleteModal.addEventListener('show.bs.modal', event =>{

        const button = event.relatedTarget;

        const id = button.getAttribute('data-bs-id');

        const inputId = document.getElementById('deleteId');

        inputId.value = id;
    });
}

document.addEventListener('DOMContentLoaded', () => {
    const deleteButton = document.getElementById('btn-excluir-contato');

    deleteButton.addEventListener('click', () => {
        const id = document.getElementById('deleteId').value;

        fetch(`/${id}`, {
            method: 'DELETE',
        })
        .then(response => {
            if(response.ok){
                alert('Contato excluído!');
                location.reload();
            }
        })
    })
})