// frontend/src/views/KundenView.jsx
import { useEffect, useState } from "react";
import { Button } from "../components/ui/Button.jsx";
import {
  ListKunden,
  CreateKunde,
} from "../../wailsjs/go/backend/App";

export default function KundenView() {
  const [kunden, setKunden] = useState([]);
  const [name, setName] = useState("");
  const [sitz, setSitz] = useState("");
  const [error, setError] = useState("");

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

        {safeKunden.length === 0 ? (
          <div className="ki-empty">Noch keine Kunden angelegt.</div>
        ) : (
          <div className="ki-list">
            {safeKunden.map((k) => (
              <div className="ki-list-item" key={k.ID}>
                <div className="ki-list-header">
                  <span className="ki-list-name">{k.Name}</span>
                  <span className="ki-list-id">#{k.ID}</span>
                </div>
                <div className="ki-list-meta">
                  <span>
                    Sitz:{" "}
                    <span className="ki-badge">
                      {k.Sitz || "—"}
                    </span>
                  </span>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
