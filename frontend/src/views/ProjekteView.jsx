// frontend/src/views/ProjekteView.jsx
import { useEffect, useState } from "react";
import { Button } from "../components/ui/Button.jsx";
import { FlexTable } from "../components/ui/FlexTable.jsx";
import {
  ListProjekte,
  CreateProjekt,
} from "../../wailsjs/go/backend/App";
import { useToasts } from "../components/ui/ToastContainer.jsx";

export default function ProjekteView() {
  const [projekte, setProjekte] = useState([]);
  const [name, setName] = useState("");
  const [kunde, setKunde] = useState([]);
  const columns = [
    { id: "id", label: "ID", field: "ID", width: 0.5, align: "center" },
    { id: "name", label: "Name", field: "Name", width: 2 },
    { id: "kunde", label: "Kunde", field: "Kunde", width: 2 },
  ];
  const { addToast } = useToasts();

  async function loadProjekte() {
    try {
      const projList = await ListProjekte();
      setProjekte(projList || []);
    } catch (e) {
      addToast({
        title: "Fehler beim Laden der Projekte",
        message: String(e),
        type: "error",
        mode: "static",
      });
      console.error(e);
    }
  }

  useEffect(() => {
    loadProjekte();
  }, []);

  async function handleSubmit(e) {
    e.preventDefault();

    try {
      await CreateProjekt({
        name,
        kunde,
      });

      setName("");
      setKunde("");
      await loadData();
    } catch (e) {
      addToast({
        title: "Fehler beim anlegen des Projektes",
        message: String(e),
        type: "error",
        mode: "static",
      });
      console.error(e);
    }
  }

  const safeProjekte = projekte || [];
  

  return (
    <div className="ki-content">
      <div className="ki-card">
        <h2 className="ki-card-title">Neues Projekt</h2>

        <form className="ki-form" onSubmit={handleSubmit}>
          <input
            className="ki-input"
            placeholder="Projektname"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
          <input
            className="ki-input"
            placeholder="Projektkunde"
            value={kunde}
            onChange={(e) => setKunde(e.target.value)}
            required
          />
          <Button onClick={handleSubmit}>Projekt anlegen</Button>
        </form>
      </div>

      <div className="ki-card">
        <h2 className="ki-card-title">Projekte</h2>
        <FlexTable columns={columns} data={safeProjekte} />
      </div>
    </div>
  );
}
