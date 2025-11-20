import { useEffect, useState } from "react";
import { ListBauteile } from "../../wailsjs/go/backend/App.js";
import { Plus } from "lucide-react";

export default function InventoryView() {
  const [bauteile, setBauteile] = useState([]);
  const [name, setName] = useState("");
  const [lagerort, setLagerort] = useState("");
  const [beschreibung, setBeschreibung] = useState("");
  const [bestand, setBestand] = useState(0);
  const [error, setError] = useState("");

  async function loadBauteile() {
    try {
      const list = await ListBauteile();
      setBauteile(list || []);
    } catch (e) {
      console.error(e);
      setError(String(e));
    }
  }

  async function handleNew() {
    // Placeholder for creating a new Bauteil
    alert("Neues Bauteil anlegen - Funktion noch nicht implementiert.");
  }

  useEffect(() => {
    loadBauteile();
  }, []);

  async function handleSubmit(e) {
    e.preventDefault();
    setError("");

    try {
      await CreateBauteil({
        name,
        lagerort,
        beschreibung,
        lagerbestand: Number(bestand),
      });

      setName("");
      setLagerort("");
      setBeschreibung("");
      setBestand(0);
      await loadBauteile();
    } catch (e) {
      console.error(e);
      setError(String(e));
    }
  }

  const safeBauteile = bauteile || [];

  return (
    <div className="ki-content">

      <div className="ki-card">
        <div className="ki-header-row">
          <h2 className="ki-card-title">Bauteile</h2>
          <button className="ki-add-btn" title="Neues Bauteil anlegen" onClick={handleNew}>
            <Plus size={16} strokeWidth={4} />
          </button>
        </div>

        {error && <div className="ki-error">{error}</div>}

        {(!bauteile || bauteile.length === 0) ? (
          <div className="ki-empty">Noch keine Bauteile angelegt.</div>
        ) : (
          <div className="ki-table">
            <div className="ki-table-head">
              <div className="ki-table-row">
                <div className="ki-table-cell ki-table-cell--head">Teil</div>
                <div className="ki-table-cell ki-table-cell--head">Sachnummer</div>
                <div className="ki-table-cell ki-table-cell--head">Kunde</div>
                <div className="ki-table-cell ki-table-cell--head">Projekt</div>
                <div className="ki-table-cell ki-table-cell--head">Erstellt</div>
              </div>
            </div>

            <div className="ki-table-body">
              {bauteile.map((b) => (
                <div className="ki-table-row" key={b.ID}>
                  <div className="ki-table-cell">{b.TeilName}</div>
                  <div className="ki-table-cell ki-mono">{b.Sachnummer}</div>
                  <div className="ki-table-cell">{b.KundeID || "—"}</div>
                  <div className="ki-table-cell">{b.ProjektID || "—"}</div>
                  <div className="ki-table-cell">
                    {formatDate(b.Erstelldatum)}
                  </div>
                </div>
              ))}
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
