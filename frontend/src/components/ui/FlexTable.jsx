import "./flexTable.css";
import { useEffect, useMemo, useState } from "react";
import { ChevronRight, ChevronLeft } from "lucide-react";
import { Modal } from "../ui/Modal.jsx";
import {ListLieferanten} from "../../../wailsjs/go/backend/App.js";


export function FlexTable({ columns, data, pageSize = 10, onUpdate}) {
  const [currentPage, setCurrentPage] = useState(1);
  const [showModal, setShowModal] = useState(false);
  const [selected, setSelected] = useState({});
  const [saving, setSaving] = useState(false);
  const [lieferanten, setLieferanten] = useState([]);

  const safeData = data || [];
  const totalItems = safeData.length;
  const totalPages = Math.max(1, Math.ceil(totalItems / pageSize));

  useEffect(() => {
    getLieferanten()
  }, [showModal]);

  async function getLieferanten() {
    const lieferanten = await ListLieferanten();
    setLieferanten(lieferanten);
  }

  useEffect(() => {
    if (currentPage > totalPages) {
      setCurrentPage(totalPages);
    }
  }, [totalPages, currentPage]);

  const pageData = useMemo(() => {
    const start = (currentPage - 1) * pageSize;
    return safeData.slice(start, start + pageSize);
  }, [safeData, currentPage, pageSize]);

  const startIndex = totalItems === 0 ? 0 : (currentPage - 1) * pageSize + 1;
  const endIndex = totalItems === 0
    ? 0
    : Math.min(currentPage * pageSize, totalItems);

  if (!columns || columns.length === 0) {
    return <div className="ki-dt-empty">Keine Spalten definiert.</div>;
  }

  const handleModalOpen = (row) => {
    setSelected(row);
    setShowModal(true);
  }

  function handleModalClose() {
    setShowModal(false);
  }

  function handleSave() {
    onUpdate(selected)
    setShowModal(false)
  }

  return (
    <div className="ki-dt-card">
      {/* Header */}
      <div className="ki-dt-header">
        {columns.map((col) => (
          <div
            key={col.id}
            className="ki-dt-cell ki-dt-cell--head"
            style={{
              flex: col.width ?? 1,
              justifyContent: mapAlign(col.align),
            }}
          >
            {col.label}
          </div>
        ))}
      </div>

      {/* Rows */}
      <div className="ki-dt-body">
        {(pageData.length === 0) ? (
          <div className="ki-dt-row ki-dt-row--empty">
            <div className="ki-dt-cell">Keine Einträge vorhanden.</div>
          </div>
        ) : (
          pageData.map((row, rowIndex) => (
            <div
              key={row.id ?? row.ID ?? rowIndex}
              className="ki-dt-row"
              onDoubleClick={() => handleModalOpen(row)}
            >
              {columns.map((col) => (
                <div
                  key={col.id}
                  className="ki-dt-cell"
                  style={{
                    flex: col.width ?? 1,
                    justifyContent: mapAlign(col.align),
                  }}
                >
                  {col.render
                    ? col.render(row)
                    : col.field
                    ? row[col.field]
                    : null}
                </div>
              ))}
            </div>
          ))
        )}
      </div>
      {/* Pagination */}
      <div className="ki-dt-footer">
        <div className="ki-dt-page-info">
          {totalItems === 0
            ? "Keine Einträge"
            : `Zeige ${startIndex}–${endIndex} von ${totalItems}`}
        </div>

        <div className="ki-dt-pagination">
          <button
            type="button"
            className="ki-dt-page-btn"
            onClick={() => setCurrentPage((p) => Math.max(1, p - 1))}
            disabled={currentPage === 1}
          >
            <ChevronLeft size={18} strokeWidth={2}/>
          </button>
          <span className="ki-dt-page-indicator">
            Seite {currentPage} / {totalPages}
          </span>
          <button
            type="button"
            className="ki-dt-page-btn"
            onClick={() =>
              setCurrentPage((p) => Math.min(totalPages, p + 1))
            }
            disabled={currentPage === totalPages || totalItems === 0}
          >
            <ChevronRight size={18} strokeWidth={2}/>
          </button>
        </div>
      </div>
      {showModal && selected ?
        <Modal
          title={"Bauteil: " + selected[columns.find(obj => obj.id === "name").field]}
          onClose={handleModalClose}
        >
          <form className="ki-form" onSubmit={() => {
          }}>
            {Object.keys(selected).map((fieldName) => {
              if (!fieldName.includes("Id") && !fieldName.includes("ID")) {
                return(
                  <div className="ki-form-group" key={fieldName}>
                    <label>{fieldName}</label>
                    {typeof selected[fieldName] === "string" ? (
                      <input
                        className="ki-input"
                        value={selected[fieldName]}
                        onChange={(e) => {
                          let obj = {...selected}
                          obj[fieldName] = e.target.value;
                          setSelected(obj);
                        }}
                        disabled={["Erstelldatum", "Sachnummer"].includes(fieldName)}
                      />
                    ) : (
                      <select
                        className="ki-input"
                        value={selected["LieferantIds"]}
                        onChange={(e) => {
                          let values = Array.from(e.target.selectedOptions).map(option => parseInt(option.value));
                          let obj = {...selected}
                          obj["LieferantIds"] = values;
                          setSelected(obj);
                        }}
                        required
                        multiple={true}
                      >
                        {lieferanten.map((lieferant) => {
                          if (fieldName === "Lieferanten") {
                            return (
                              <option key={lieferant.ID} value={lieferant.ID}>
                                {lieferant.Name} - {lieferant.Sitz}
                              </option>
                            )
                          }

                        })}
                      </select>
                    )}

                  </div>
                )
              }
            })}
            <div className="ki-form-actions">
              <button
                type="button"
                className="ki-btn-secondary"
                onClick={handleModalClose}
                disabled={saving}
              >
                Abbrechen
              </button>
              <button
                type="submit"
                className="ki-btn-primary"
                onClick={handleSave}
                disabled={saving}
              >
                {saving ? "Speichern…" : "Bauteil aktualisieren"}
              </button>
            </div>
          </form>
        </Modal>
      : ""}
    </div>
  );
}

function mapAlign(align) {
  switch (align) {
    case "center":
      return "center";
    case "right":
      return "flex-end";
    case "left":
    default:
      return "flex-start";
  }
}
