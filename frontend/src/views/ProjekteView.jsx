// frontend/src/views/ProjekteView.jsx
import { useEffect, useState } from "react";
import { Button } from "../components/ui/Button.jsx";
import {
  ListProjekte,
  CreateProjekt,
  ListKunden,
} from "../../wailsjs/go/backend/App";

export default function ProjekteView() {
  const [projekte, setProjekte] = useState([]);
  const [kunden, setKunden] = useState([]);
  const [name, setName] = useState("");
  const [kundeID, setKundeID] = useState("");
  const [error, setError] = useState("");

  async function loadData() {
    try {
      const [projList, kundenList] = await Promise.all([
        ListProjekte(),
        ListKunden(),
      ]);
      setProjekte(projList || []);
      setKunden(kundenList || []);
    } catch (e) {
      console.error(e);
      setError(String(e));
    }
  }

  useEffect(() => {
    loadData();
  }, []);

  async function handleSubmit(e) {
    e.preventDefault();
    setError("");

    if (!kundeID) {
      setError("Bitte einen Kunden auswählen.");
      return;
    }

    try {
      await CreateProjekt({
        name,
        kundeID: Number(kundeID),
      });

      setName("");
      setKundeID("");
      await loadData();
    } catch (e) {
      console.error(e);
      setError(String(e));
    }
  }

  const safeProjekte = projekte || [];
  const safeKunden = kunden || [];

  function getKundenName(id) {
    const k = safeKunden.find((c) => c.ID === id);
    return k ? k.Name : `Kunde #${id}`;
  }

  return (
    <div className="ki-content">
      <div className="ki-card">
        <h2 className="ki-card-title">Neues Projekt</h2>
        {error && <div className="ki-error">{error}</div>}

        <form className="ki-form" onSubmit={handleSubmit}>
          <input
            className="ki-input"
            placeholder="Projektname"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />

          <select
            className="ki-input"
            value={kundeID}
            onChange={(e) => setKundeID(e.target.value)}
            required
          >
            <option value="">Kunde auswählen…</option>
            {safeKunden.map((k) => (
              <option key={k.ID} value={k.ID}>
                {k.Name} {k.Sitz ? `(${k.Sitz})` : ""}
              </option>
            ))}
          </select>

          <Button onClick={handleSubmit}>Projekt anlegen</Button>
        </form>
      </div>

      <div className="ki-card">
        <h2 className="ki-card-title">Projekte</h2>

        {safeProjekte.length === 0 ? (
          <div className="ki-empty">Noch keine Projekte angelegt.</div>
        ) : (
          <div className="ki-list">
            {safeProjekte.map((p) => (
              <div className="ki-list-item" key={p.ID}>
                <div className="ki-list-header">
                  <span className="ki-list-name">{p.Name}</span>
                  <span className="ki-list-id">#{p.ID}</span>
                </div>

                <div className="ki-list-meta">
                  <span>
                    Kunde:{" "}
                    <span className="ki-badge">
                      {getKundenName(p.KundeID)}
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
