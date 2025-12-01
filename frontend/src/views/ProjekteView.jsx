import { useEffect, useState } from "react";
import { FlexTable } from "../components/ui/FlexTable.jsx";
import {NewProjektModal} from "../components/special/NewProjektModal.jsx";
import { Plus } from "lucide-react";
import {
  ListProjekte,
  CreateProjekt,
} from "../../wailsjs/go/backend/App";
import { useToasts } from "../components/ui/ToastContainer.jsx";

export default function ProjekteView() {
  const [projekte, setProjekte] = useState([]);
  const [modalOpen, setModalOpen] = useState(false);

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

  function handleProjektCreated() {
    loadProjekte()
  }

  const safeProjekte = projekte || [];

  return (
    <div className="ki-content">
      <div className="ki-card">
        <div className="ki-header-row">
          <h2 className="ki-card-title">Projekte</h2>
          <button className="ki-add-btn" title="Neues Projekt anlegen" onClick={() => setModalOpen(true)}>
            <Plus size={16} strokeWidth={4} />
          </button>
        </div>
        <FlexTable columns={columns} data={safeProjekte} />
      </div>
      <NewProjektModal
        open={modalOpen}
        onClose={() => setModalOpen(false)}
        onCreated={() => {
          handleProjektCreated();
        }}
      />
    </div>
  );
}
