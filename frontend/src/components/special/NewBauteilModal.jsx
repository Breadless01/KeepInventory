import { useEffect, useMemo, useState } from "react";
import { Modal } from "../ui/Modal.jsx";
import { Stepper } from "../ui/Stepper.jsx";
import "./newBauteilModal.css"

import {
  CreateBauteil,
  ListTypen,
  ListHerstellungsarten,
  ListVerschleissteile,
  ListFunktionen,
  ListMaterialien,
  ListOberflaechenbehandlungen,
  ListFarben,
  ListReserven,
} from "../../../wailsjs/go/backend/App";


export function NewBauteilModal({ open, onClose, onCreated }) {
  const [loading, setLoading] = useState(false);
  const [saving, setSaving] = useState(false);
  const [error, setError] = useState("");

  // Stammdaten
  const [typen, setTypen] = useState([]);
  const [arten, setArten] = useState([]);
  const [verschleissteile, setVerschleissteile] = useState([]);
  const [funktionen, setFunktionen] = useState([]);
  const [materialien, setMaterialien] = useState([]);
  const [oberflaechen, setOberflaechen] = useState([]);
  const [farben, setFarben] = useState([]);
  const [reserven, setReserven] = useState([]);

  // Formularwerte
  const [teilName, setTeilName] = useState("");
  const [kundeID, setKundeID] = useState("");
  const [projektID, setProjektID] = useState("");

  const [typID, setTypID] = useState("");
  const [artID, setArtID] = useState("");
  const [verschID, setVerschID] = useState("");
  const [funktionID, setFunktionID] = useState("");
  const [materialID, setMaterialID] = useState("");
  const [oberfID, setOberfID] = useState("");
  const [farbeID, setFarbeID] = useState("");
  const [reserveID, setReserveID] = useState("");
  const [steps, setSteps] = useState([
      {
        id: 1,
        label: "Basisdaten",
        content: <div>Hier kommen die Basisdaten rein.</div>,
      },
      {
        id: 2,
        label: "Merkmale",
        content: <div>Hier kommen die Merkmale rein.</div>,
      },
      {
        id: 3,
        label: "Details",
        content: <div>Hier kommen die Details rein.</div>,
      },
    ]);
    const [currentIndex, setCurrentIndex] = useState(0);
    const [nextId, setNextId] = useState(3);


  // Stammdaten laden, wenn Modal geöffnet wird
  useEffect(() => {
    if (!open) return;

    async function loadStammdaten() {
      try {
        setLoading(true);
        setError("");

        const [
          t,
          a,
          v,
          f,
          m,
          o,
          c,
          r,
        ] = await Promise.all([
          ListTypen(),
          ListHerstellungsarten(),
          ListVerschleissteile(),
          ListFunktionen(),
          ListMaterialien(),
          ListOberflaechenbehandlungen(),
          ListFarben(),
          ListReserven(),
        ]);

        setTypen(t || []);
        setArten(a || []);
        setVerschleissteile(v || []);
        setFunktionen(f || []);
        setMaterialien(m || []);
        setOberflaechen(o || []);
        setFarben(c || []);
        setReserven(r || []);
      } catch (e) {
        console.error(e);
        setError(String(e));
      } finally {
        setLoading(false);
      }
    }

    loadStammdaten();
  }, [open]);

  // Preview der Sachnummer – Struktur + Platzhalter für Hex-Suffix
  const sachnummerPreview = useMemo(() => {
    if (
      !typID ||
      !artID ||
      !verschID ||
      !funktionID ||
      !materialID ||
      !oberfID ||
      !farbeID ||
      !reserveID
    ) {
      return "Bitte alle Merkmalsfelder wählen, um eine Vorschau zu sehen.";
    }

    const findById = (xs, id) => xs.find((x) => x.ID === Number(id));

    const typ = findById(typen, typID);
    const art = findById(arten, artID);
    const versch = findById(verschleissteile, verschID);
    const fun = findById(funktionen, funktionID);
    const mat = findById(materialien, materialID);
    const oberf = findById(oberflaechen, oberfID);
    const farbe = findById(farben, farbeID);
    const res = findById(reserven, reserveID);

    if (!typ || !art || !versch || !fun || !mat || !oberf || !farbe || !res) {
      return "Auswahl unvollständig oder Stammdaten nicht geladen.";
    }

    // Wir nutzen die Symbol-Werte, so wie im Backend
    const keyPart = `${typ.Symbol}-${art.Symbol}-${versch.Symbol}-${fun.Symbol}-${mat.Symbol}-${oberf.Symbol}-${farbe.Symbol}-${res.Symbol}`;

    // Hex-Suffix wird im Backend final aus der DB berechnet.
    // Hier zeigen wir nur die Struktur + Platzhalter.
    const fakeSuffix = "????";

    return `${keyPart}-${fakeSuffix}`;
  }, [
    typID,
    artID,
    verschID,
    funktionID,
    materialID,
    oberfID,
    farbeID,
    reserveID,
    typen,
    arten,
    verschleissteile,
    funktionen,
    materialien,
    oberflaechen,
    farben,
    reserven,
  ]);

  async function handleSubmit(e) {
    e?.preventDefault();
    setError("");

    // minimale Validierung
    if (!teilName.trim()) {
      setError("Bitte einen Teilnamen eingeben.");
      return;
    }
    if (
      !typID ||
      !artID ||
      !verschID ||
      !funktionID ||
      !materialID ||
      !oberfID ||
      !farbeID ||
      !reserveID
    ) {
      setError("Bitte alle Merkmalsfelder auswählen.");
      return;
    }

    try {
      setSaving(true);

      const req = {
        TeilName: teilName,
        KundeID: kundeID,
        ProjektID: projektID,

        TypID: Number(typID),
        HerstellungsartID: Number(artID),
        VerschleissteilID: Number(verschID),
        FunktionID: Number(funktionID),
        MaterialID: Number(materialID),
        OberflaechenbehandlungID: Number(oberfID),
        FarbeID: Number(farbeID),
        ReserveID: Number(reserveID),
      };

      const created = await CreateBauteil(req);

      if (onCreated) {
        onCreated(created);
      }

      // Formular zurücksetzen
      setTeilName("");
      setKundeID("");
      setProjektID("");
      setTypID("");
      setArtID("");
      setVerschID("");
      setFunktionID("");
      setMaterialID("");
      setOberfID("");
      setFarbeID("");
      setReserveID("");

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
    <Modal title="Neues Bauteil" onClose={onClose}>
      {loading ? (
        <div>Stammdaten werden geladen…</div>
      ) : (
        <form className="ki-form" onSubmit={handleSubmit}>
          <Stepper
            steps={steps}
            currentIndex={currentIndex}
            onStepChange={setCurrentIndex}
            showLabels={true}
          />
          {error && <div className="ki-error">{error}</div>}

          <div className="ki-form-group">
            <label>Teil-Name</label>
            <input
              className="ki-input"
              value={teilName}
              onChange={(e) => setTeilName(e.target.value)}
              required
            />
          </div>

          <div className="ki-form-row">
            <div className="ki-form-group">
              <label>Kunde-ID</label>
              <input
                className="ki-input"
                value={kundeID}
                onChange={(e) => setKundeID(e.target.value)}
                placeholder="optional"
              />
            </div>
            <div className="ki-form-group">
              <label>Projekt-ID</label>
              <input
                className="ki-input"
                value={projektID}
                onChange={(e) => setProjektID(e.target.value)}
                placeholder="optional"
              />
            </div>
          </div>

          <hr className="ki-separator" />

          <div className="ki-form-row">
            <div className="ki-form-group">
              <label>Typ</label>
              <select
                className="ki-input"
                value={typID}
                onChange={(e) => setTypID(e.target.value)}
                required
              >
                <option value="">— wählen —</option>
                {typen.map((t) => (
                    <option key={t.ID} value={t.ID}>
                        {t.Symbol} – {t.Name}
                    </option>
                ))}
              </select>
            </div>

            <div className="ki-form-group">
              <label>Herstellungsart</label>
              <select
                className="ki-input"
                value={artID}
                onChange={(e) => setArtID(e.target.value)}
                required
              >
                <option value="">— wählen —</option>
                {arten.map((a) => (
                  <option key={a.ID} value={a.ID}>
                    {a.Symbol} – {a.Name}
                  </option>
                ))}
              </select>
            </div>
          </div>

          <div className="ki-form-row">
            <div className="ki-form-group">
              <label>Verschleißteil</label>
              <select
                className="ki-input"
                value={verschID}
                onChange={(e) => setVerschID(e.target.value)}
                required
              >
                <option value="">— wählen —</option>
                {verschleissteile.map((v) => (
                  <option key={v.ID} value={v.ID}>
                    {v.Symbol} – {v.Name}
                  </option>
                ))}
              </select>
            </div>

            <div className="ki-form-group">
              <label>Funktion</label>
              <select
                className="ki-input"
                value={funktionID}
                onChange={(e) => setFunktionID(e.target.value)}
                required
              >
                <option value="">— wählen —</option>
                {funktionen.map((f) => (
                  <option key={f.ID} value={f.ID}>
                    {f.Symbol} – {f.Name}
                  </option>
                ))}
              </select>
            </div>
          </div>

          <div className="ki-form-row">
            <div className="ki-form-group">
              <label>Material</label>
              <select
                className="ki-input"
                value={materialID}
                onChange={(e) => setMaterialID(e.target.value)}
                required
              >
                <option value="">— wählen —</option>
                {materialien.map((m) => (
                  <option key={m.ID} value={m.ID}>
                    {m.Symbol} – {m.Name}
                  </option>
                ))}
              </select>
            </div>

            <div className="ki-form-group">
              <label>Oberflächenbeh.</label>
              <select
                className="ki-input"
                value={oberfID}
                onChange={(e) => setOberfID(e.target.value)}
                required
              >
                <option value="">— wählen —</option>
                {oberflaechen.map((o) => (
                  <option key={o.ID} value={o.ID}>
                    {o.Symbol} – {o.Name}
                  </option>
                ))}
              </select>
            </div>
          </div>

          <div className="ki-form-row">
            <div className="ki-form-group">
              <label>Farbe</label>
              <select
                className="ki-input"
                value={farbeID}
                onChange={(e) => setFarbeID(e.target.value)}
                required
              >
                <option value="">— wählen —</option>
                {farben.map((c) => (
                  <option key={c.ID} value={c.ID}>
                    {c.Symbol} – {c.Name}
                  </option>
                ))}
              </select>
            </div>

            <div className="ki-form-group">
              <label>Reserve</label>
              <select
                className="ki-input"
                value={reserveID}
                onChange={(e) => setReserveID(e.target.value)}
                required
              >
                <option value="">— wählen —</option>
                {reserven.map((r) => (
                  <option key={r.ID} value={r.ID}>
                    {r.Symbol} – {r.Name}
                  </option>
                ))}
              </select>
            </div>
          </div>

          <div className="ki-form-group">
            <label>Sachnummer (Vorschau)</label>
            <div className="ki-sachnummer-preview">
              {sachnummerPreview}
            </div>
            <div className="ki-sachnummer-hint">
              Finale Sachnummer (inkl. Hex-Suffix) wird im Backend vergeben.
            </div>
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
      )}
    </Modal>
  );
}