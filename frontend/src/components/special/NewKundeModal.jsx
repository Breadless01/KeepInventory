import { useEffect, useState } from "react";
import { Modal } from "../ui/Modal.jsx";
import "./newModal.css"

import {
  CreateKunde
} from "../../../wailsjs/go/backend/App";


export function NewKundeModal({ open, onClose, onCreated }) {
  const [saving, setSaving] = useState(false);
  const [error, setError] = useState("");

  // Formularwerte
  const [name, setName] = useState("");
  const [sitz, setSitz] = useState("");

  useEffect(() => {
    if (!open) return;
    setError("");
  }, [open])

  async function handleSubmit(e) {
    e?.preventDefault();
    setError("");

    if (!name.trim()) {
      setError("Bitte einen Namen eingeben.");
      return;
    }

    if (!sitz.trim()) {
      setError("Bitte einen Sitz eingeben.");
      return;
    }

    try {
      setSaving(true);

      const req = {
        Name: name,
        Sitz: sitz
      };

      const created = await CreateKunde(req);

      if (onCreated) {
        onCreated(created);
      }

      // Formular zurücksetzen
      setName("");
      setSitz("");

      onClose();
    } catch (e) {
      console.error(e);
      setError(String(e));
    } finally {
      setSaving(false);
    }
  }

  if (!open) return null;

  return (
    <Modal title="Neuer Kunde" onClose={onClose}>
      <form className="ki-form" onSubmit={handleSubmit}>
        {error && <div className="ki-error">{error}</div>}
        <div className="ki-form-group">
          <label>Name*</label>
          <input
            className="ki-input"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
        </div>
        <div className="ki-form-group">
          <label>Sitz*</label>
          <input
            className="ki-input"
            value={sitz}
            onChange={(e) => setSitz(e.target.value)}
            required
        />
        </div>
          <div className="ki-form-actions">
            <button
              type="button"
              className="ki-btn-secondary"
              onClick={onClose}
              disabled={saving}
            >
              Abbrechen
            </button>
            <button
              type="submit"
              className="ki-btn-primary"
              disabled={saving}
            >
              {saving ? "Speichern…" : "Bauteil anlegen"}
            </button>
          </div>
      </form>
    </Modal>
  );
}