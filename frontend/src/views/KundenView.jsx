// frontend/src/views/KundenView.jsx
import { useEffect, useState } from "react";
import { Button } from "../components/ui/Button.jsx";
import { FlexTable } from "../components/ui/FlexTable.jsx";

import {
  ListKunden,
  CreateKunde,
} from "../../wailsjs/go/backend/App";

export default function KundenView() {
  const [kunden, setKunden] = useState([]);
  const [name, setName] = useState("");
  const [sitz, setSitz] = useState("");
  const [error, setError] = useState("");
  const columns = [
    { id: "id", label: "ID", field: "ID", width: 0.5, align: "center" },
    { id: "name", label: "Name", field: "Name", width: 2 },
    { id: "sitz", label: "Sitz", field: "Sitz", width: 2 },
  ];

  async function loadKunden() {
    try {
      const list = await ListKunden();
      setKunden(list || []);
    } catch (e) {
      console.error(e);
      setError(String(e));
    }
  }

  useEffect(() => {
    loadKunden();
  }, []);

  async function handleSubmit(e) {
    e.preventDefault();
    setError("");

    try {
      await CreateKunde({
        name,
        sitz,
      });

      setName("");
      setSitz("");
      await loadKunden();
    } catch (e) {
      console.error(e);
      setError(String(e));
    }
  }

  const safeKunden = kunden || [];

  return (
    <div className="ki-content">
      <div className="ki-card">
        <h2 className="ki-card-title">Neuer Kunde</h2>
        {error && <div className="ki-error">{error}</div>}

        <form className="ki-form" onSubmit={handleSubmit}>
          <input
            className="ki-input"
            placeholder="Name des Kunden"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
          <input
            className="ki-input"
            placeholder="Sitz (Ort / Adresse)"
            value={sitz}
            onChange={(e) => setSitz(e.target.value)}
          />

          <Button onClick={handleSubmit}>Kunden anlegen</Button>
        </form>
      </div>

      <div className="ki-card">
        <h2 className="ki-card-title">Kunden</h2>
        <FlexTable columns={columns} data={safeKunden} />
      </div>
    </div>
  );
}
