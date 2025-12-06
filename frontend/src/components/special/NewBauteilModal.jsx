import { useEffect, useMemo, useState } from "react";
import { Modal } from "../ui/Modal.jsx";
import { Stepper } from "../ui/Stepper.jsx";
import "./NewModal.css"

import {
  CreateBauteil,
  ListLieferanten,
  ListTypen,
  ListHerstellungsarten,
  ListVerschleissteile,
  ListFunktionen,
  ListMaterialien,
  ListOberflaechenbehandlungen,
  ListFarben,
  ListReserven,
  ListKunden,
  ListProjekte
} from "../../../wailsjs/go/backend/App";
import MultiSelectCheckboxGroup from "../ui/MultiSelectCheckboxGroup.jsx";


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

  // Kunden ,Projekte, Lieferanten
  const [kunden, setKunden] = useState([]);
  const [projekte, setProjekte] = useState([]);
  const [lieferanten, setLieferanten] = useState([]);

  // Formularwerte
  const [teilName, setTeilName] = useState("");
  const [kundeID, setKundeID] = useState("");
  const [projektID, setProjektID] = useState("");
  const [lieferantenIds, setLieferantenIds] = useState([])

  const [typID, setTypID] = useState("");
  const [artID, setArtID] = useState("");
  const [verschID, setVerschID] = useState("");
  const [funktionID, setFunktionID] = useState("");
  const [materialID, setMaterialID] = useState("");
  const [oberfID, setOberfID] = useState("");
  const [farbeID, setFarbeID] = useState("");
  const [reserveID, setReserveID] = useState("");
  const [currentStep, setCurrentStep] = useState(0);


  useEffect(() => {
    if (!open) return;

    setCurrentStep(0);
    setError("");

    async function loadDaten() {
      try {
        setLoading(true);
        const [
          t,
          a,
          v,
          f,
          m,
          o,
          c,
          r,
          kunde,
          projekt,
          lieferanten
        ] = await Promise.all([
          ListTypen(),
          ListHerstellungsarten(),
          ListVerschleissteile(),
          ListFunktionen(),
          ListMaterialien(),
          ListOberflaechenbehandlungen(),
          ListFarben(),
          ListReserven(),
          ListKunden(),
          ListProjekte(),
          ListLieferanten(),
        ]);

        setTypen(t || []);
        setArten(a || []);
        setVerschleissteile(v || []);
        setFunktionen(f || []);
        setMaterialien(m || []);
        setOberflaechen(o || []);
        setFarben(c || []);
        setReserven(r || []);
        setKunden(kunde || []);
        setProjekte(projekt || []);
        setLieferanten(lieferanten || [])
      } catch (e) {
        console.error(e);
        setError(String(e));
      } finally {
        setLoading(false);
      }
    }

    loadDaten();
  }, [open]);

  // Sachnummer-Preview auf Basis der ausgewählten Symbole
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

    const keyPart = `${typ.Symbol}${art.Symbol}${versch.Symbol}${fun.Symbol}${mat.Symbol}${oberf.Symbol}${farbe.Symbol}${res.Symbol}`;

    // echtes Hex-Suffix kommt aus dem Backend → hier nur Dummy/Struktur
    const fakeSuffix = "????";

    return `${keyPart}${fakeSuffix}`;
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

    if (!teilName.trim()) {
      setError("Bitte einen Teilnamen eingeben.");
      setCurrentStep(0);
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
      setCurrentStep(1);
      return;
    }

    try {
      setSaving(true);
      console.log(lieferantenIds)

      const req = {
        TeilName: teilName,
        KundeID: kundeID ? Number(kundeID) : undefined,
        ProjektID: projektID ? Number(projektID) : undefined,
        LieferantenIds: lieferantenIds,
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

  function goNextStep() {
    if (currentStep === 0) {
      if (!teilName.trim()) {
        setError("Bitte einen Teilnamen eingeben.");
        return;
      }
    }
    if (currentStep === 1) {
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
    }
    setError("");
    setCurrentStep((s) => Math.min(s + 1, 2));
  }

  function goPrevStep() {
    setError("");
    setCurrentStep((s) => Math.max(s - 1, 0));
  }

  if (!open) return null;

  // Step-Inhalte definieren
  const steps = [
    {
      id: "basis",
      label: "Basisdaten",
      content: (
        <>
          <div className="ki-form-group">
            <label>Teil-Name*</label>
            <input
              className="ki-input"
              value={teilName}
              onChange={(e) => setTeilName(e.target.value)}
              required
            />
          </div>

          <div className="ki-form-group">
            <label>Kunde</label>
            <select
              className="ki-input"
              value={kundeID}
              onChange={(e) => setKundeID(e.target.value)}
              required
            >
              <option value="">— wählen —</option>
              {kunden.map((kunde) => (
                <option key={kunde.ID} value={kunde.ID}>
                  {kunde.Name} {kunde.Sitz && (`– ${kunde.Sitz}`)}
                </option>
              ))}
            </select>
          </div>
          <div className="ki-form-group">
            <label>Projekt</label>
            <select
              className="ki-input"
              value={projektID}
              onChange={(e) => setProjektID(e.target.value)}
              required
            >
              <option value="">— wählen —</option>
              {projekte.map((projekt) => (
                <option key={projekt.ID} value={projekt.ID}>
                  {projekt.Name} {projekt.Kunde && (`– ${projekt.Kunde}`)}
                </option>
              ))}
            </select>
          </div>
          <div className="ki-form-group">
            <label>Lieferanten</label>
            <select
              className="ki-input"
              value={lieferantenIds}
              onChange={(e) => {
                setLieferantenIds([...lieferantenIds, parseInt(e.target.value)])
              }}
              required
              multiple={true}
            >
              {lieferanten.map((lieferant) => (
                <option key={lieferant.ID} value={lieferant.ID}>
                  {lieferant.Name}
                </option>
              ))}
            </select>
          </div>
        </>
      ),
    },
    {
      id: "merkmale",
      label: "Merkmale",
      content: (
        <>
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
        </>
      ),
    },
    {
      id: "sachnummer",
      label: "Sachnummer",
      content: (
        <>
          <div className="ki-form-group">
            <label>Sachnummer (Vorschau)</label>
            <div className="ki-sachnummer-preview">{sachnummerPreview}</div>
            <div className="ki-sachnummer-hint">
              Die finale Sachnummer (inkl. Hex-Suffix) wird beim Speichern im
              Backend berechnet.
            </div>
          </div>
        </>
      ),
    },
  ];

  return (
    <Modal title="Neues Bauteil" onClose={onClose}>
      {loading ? (
        <div>Stammdaten werden geladen…</div>
      ) : (
        <form className="ki-form" onSubmit={handleSubmit}>
          {error && <div className="ki-error">{error}</div>}

          <Stepper
            steps={steps}
            currentIndex={currentStep}
            onStepChange={setCurrentStep}
            showLabels={false}
          />

          <div className="ki-form-actions">
            <button
              type="button"
              className="ki-btn-secondary"
              onClick={onClose}
              disabled={saving}
            >
              Abbrechen
            </button>

            {currentStep > 0 && (
              <button
                type="button"
                className="ki-btn-secondary"
                onClick={goPrevStep}
                disabled={saving}
              >
                Zurück
              </button>
            )}

            {currentStep < steps.length - 1 ? (
              <button
                type="button"
                className="ki-btn-primary"
                onClick={goNextStep}
                disabled={saving}
              >
                Weiter
              </button>
            ) : (
              <button
                type="submit"
                className="ki-btn-primary"
                disabled={saving}
              >
                {saving ? "Speichern…" : "Bauteil anlegen"}
              </button>
            )}
          </div>
        </form>
      )}
    </Modal>
  );
}