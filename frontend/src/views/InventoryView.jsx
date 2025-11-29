import { useEffect, useState } from "react";
import { ListBauteile } from "../../wailsjs/go/backend/App.js";
import { NewBauteilModal } from "../components/special/NewBauteilModal.jsx";
import { Plus } from "lucide-react";
import { FlexTable } from "../components/ui/FlexTable.jsx";
import { useToasts } from "../components/ui/ToastContainer.jsx";


export default function InventoryView() {
  const [bauteile, setBauteile] = useState([]);
  const [name, setName] = useState("");
  const [lagerort, setLagerort] = useState("");
  const [beschreibung, setBeschreibung] = useState("");
  const [bestand, setBestand] = useState(0);
  const [modalOpen, setModalOpen] = useState(false);
  const columns = [
    { id: "id", label: "ID", field: "ID", width: 0.5, align: "center" },
    { id: "name", label: "TeilName", field: "TeilName", width: 2 },
    { id: "sachnummer", label: "Sachnummer", field: "Sachnummer", width: 2 },
    { id: "kunde", label: "Kunde", field: "Kunde", width: 2 },
    { id: "projekt", label: "Projekt", field: "Projekt", width: 2 },
    { id: "erstelldatum", label: "Erstelldatum", field: "Erstelldatum", width: 2 },
  ];

  const { addToast } = useToasts();
  
  async function loadBauteile() {
    try {
      const list = await ListBauteile();
      setBauteile(list || []);
    } catch (e) {
      addToast({
        title: "Fehler beim Laden der Bauteile",
        message: String(e),
        type: "error",
        mode: "static",
      });
      console.error(e);
    }
  }

  useEffect(() => {
    loadBauteile();
  }, []);

  async function reloadBauteile() {
    try {
      const list = await ListBauteile();
      setBauteile(list || []);
    } catch (e) {
      addToast({
        title: "Fehler beim Laden der Bauteile",
        message: String(e),
        type: "error",
        mode: "static",
      });
      console.error(e);
    }
  }

  useEffect(() => {
    reloadBauteile();
  }, []);

  const safeBauteile = bauteile || [];
  console.log(safeBauteile);
  

  return (
    <div className="ki-content">

      <div className="ki-card">
        <div className="ki-header-row">
          <h2 className="ki-card-title">Bauteile</h2>
          <button className="ki-add-btn" title="Neues Bauteil anlegen" onClick={() => setModalOpen(true)}>
            <Plus size={16} strokeWidth={4} />
          </button>
        </div>

        <FlexTable columns={columns} data={safeBauteile} />
      </div>
      <NewBauteilModal
        open={modalOpen}
        onClose={() => setModalOpen(false)}
        onCreated={() => {
          reloadBauteile();
        }}
      />
    </div>
  );
}
