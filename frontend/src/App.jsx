import React, { useEffect, useState } from 'react';

function App() {
  const [items, setItems] = useState([]);
  const [name, setName] = useState('');
  const [value, setValue] = useState('');
  const [editId, setEditId] = useState(null);

  useEffect(() => {
    fetch('/items')
      .then(res => res.json())
      .then(setItems);
  }, []);

  const handleSubmit = (e) => {
    e.preventDefault();
    if (editId) {
      fetch(`/item?id=${editId}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name, value })
      })
        .then(() => reload());
    } else {
      fetch('/items', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name, value })
      })
        .then(() => reload());
    }
    setName('');
    setValue('');
    setEditId(null);
  };

  const reload = () => {
    fetch('/items')
      .then(res => res.json())
      .then(setItems);
  };

  const handleEdit = (item) => {
    setEditId(item.id);
    setName(item.name);
    setValue(item.value);
  };

  const handleDelete = (id) => {
    fetch(`/item?id=${id}`, { method: 'DELETE' })
      .then(() => reload());
  };

  return (
    <div style={{ maxWidth: 600, margin: '2rem auto', fontFamily: 'sans-serif' }}>
      <h1>Go CRUD Items</h1>
      <form onSubmit={handleSubmit} style={{ marginBottom: 20 }}>
        <input
          placeholder="Name"
          value={name}
          onChange={e => setName(e.target.value)}
          required
        />
        <input
          placeholder="Value"
          value={value}
          onChange={e => setValue(e.target.value)}
          required
        />
        <button type="submit">{editId ? 'Update' : 'Add'}</button>
        {editId && <button type="button" onClick={() => { setEditId(null); setName(''); setValue(''); }}>Cancel</button>}
      </form>
      <ul>
        {items.map(item => (
          <li key={item.id} style={{ marginBottom: 10 }}>
            <b>{item.name}</b>: {item.value}
            <button onClick={() => handleEdit(item)} style={{ marginLeft: 10 }}>Edit</button>
            <button onClick={() => handleDelete(item.id)} style={{ marginLeft: 5 }}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default App;
