import { useEffect, useState } from "react";
import { GetFilterConfig, FilterBauteile, UpdateBauteil } from "../../wailsjs/go/backend/App.js";
import { NewBauteilModal } from "../components/special/NewBauteilModal.jsx";
import FacetFilterPanel from "../components/ui/FacetFilterPanel.jsx"
import { Plus } from "lucide-react";
import { FlexTable } from "../components/ui/FlexTable.jsx";
import { useToasts } from "../components/ui/ToastContainer.jsx";
import { Searchbar } from "../components/ui/Searchbar.jsx";

export default function InventoryView() {
  const [bauteile, setBauteile] = useState([]);
  const [modalOpen, setModalOpen] = useState(false);

  const columns = [
    { id: "id", label: "ID", field: "ID", width: 0.5, align: "center" },
    { id: "name", label: "TeilName", field: "TeilName", width: 2 },
    { id: "sachnummer", label: "Sachnummer", field: "Sachnummer", width: 2 },
    { id: "kunde", label: "Kunde", field: "Kunde", width: 2 },
    { id: "projekt", label: "Projekt", field: "Projekt", width: 2 },
    { id: "erstelldatum", label: "Erstelldatum", field: "Erstelldatum", width: 2 },
  ];

  // Filter-spezifischer State
  const [filterConfig, setFilterConfig] = useState(null);
  const [filterState, setFilterState] = useState({});
  const [filterIds, setFilterIds] = useState([]);
  const [facets, setFacets] = useState({});
  const [total, setTotal] = useState(0);
  const [page, setPage] = useState(1);
  const pageSize = 20;

  const { addToast } = useToasts();

  useEffect(() => {
    async function init() {
      try {
        const cfg = await GetFilterConfig();
        setFilterConfig(cfg);

        await applyFilter(cfg, {}, 1);
      } catch (e) {
        addToast({
          title: "Fehler beim Laden der Filter-Konfiguration",
          message: String(e),
          type: "error",
          mode: "static",
        });
        console.error(e);
      }
    }

    init();
  }, []);

  async function applyFilter(cfg, state, pageNumber) {
    try {
      const res = await FilterBauteile({
        page: pageNumber,
        pageSize,
        filters: state,
      });

      setBauteile(res.items || []);
      setFacets(res.facets || {});
      setTotal(res.total || 0);
      setPage(pageNumber);
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
    setFilterState({ ...filterState, ["id"]: filterIds })
  }, [filterIds]);

  useEffect(() => {
    if (!filterConfig) return;
    // setFilterState({ ...filterState, ["id"]: filterIds })
    applyFilter(filterConfig, filterState, 1);
  }, [filterState, filterConfig]);

  function handleBauteilCreated() {
    if (!filterConfig) {
      applyFilter({ resources: [] }, {}, 1);
    } else {
      applyFilter(filterConfig, filterState, page);
    }
  }

  async function handleBauteilUpdate(bauteil) {
    const req = {
      ID: bauteil.ID,
      TeilName: bauteil.TeilName,
      KundeId: bauteil.KundeId,
      ProjektId: bauteil.ProjektId,
      LieferantenIds: bauteil.LieferantIds
    };
    const res = await UpdateBauteil(req)
    if (!filterConfig) {
      applyFilter({ resources: [] }, {}, 1);
    } else {
      applyFilter(filterConfig, filterState, page);
    }
  }

  const safeBauteile = bauteile || [];

  let facetFieldConfigs = [];
  if (filterConfig && Array.isArray(filterConfig.resources)) {
    const bauteilRes = filterConfig.resources.find(
      (r) => r.resource === "bauteile"
    );
    if (bauteilRes && Array.isArray(bauteilRes.fields)) {
      facetFieldConfigs = bauteilRes.fields.filter((f) => f.enabled);
    }
  }
  return (
    <div className="ki-content">
      <Searchbar
        objType="bauteil"
        onEnter={setFilterIds}
      />
      <div className="ki-card">
        <div className="ki-header-row">
          <h2 className="ki-card-title">Bauteile</h2>
          <button className="ki-add-btn" title="Neues Bauteil anlegen" onClick={() => setModalOpen(true)}>
            <Plus size={16} strokeWidth={4} />
          </button>
        </div>
        {facetFieldConfigs.length > 0 && (
          <div style={{ marginBottom: "1rem" }}>
            <FacetFilterPanel
              total={total}
              facets={facets}
              filterState={filterState}
              onChange={setFilterState}
              fieldConfigs={facetFieldConfigs}
            />
          </div>
        )}
        <FlexTable columns={columns} data={safeBauteile} onUpdate={(bauteil) => {
          handleBauteilUpdate(bauteil)
        }}/>
      </div>
      <NewBauteilModal
        open={modalOpen}
        onClose={() => setModalOpen(false)}
        onCreated={() => {
          handleBauteilCreated();
        }}
      />
    </div>
  );
}
