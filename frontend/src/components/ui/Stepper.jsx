import "./stepper.css";
import { Check } from "lucide-react"

export function Stepper({
  steps,
  currentIndex,
  showLabels = true,
  onStepChange,
  onAddStep,
}) {
  const hasAdd = typeof onAddStep === "function";

  const maxIndex = steps.length > 0 ? steps.length - 1 : 0;
  const progress =
    maxIndex > 0 ? Math.min(100, Math.max(0, (currentIndex / maxIndex) * 100)) : 0;


  return (
    <div className="ki-stepper">
      <div className="ki-stepper-header">
        <div className="ki-stepper-steps-wrapper">
            <div className="ki-stepper-track">
                <div
                className="ki-stepper-track-fill"
                style={{ width: `${progress}%` }}
                />
            </div>
            <div className="ki-stepper-steps">
            {steps.map((step, index) => {
                const isActive = index === currentIndex;
                const isDone = index < currentIndex;

                return (
                <button
                    key={step.id ?? index}
                    className={
                    "ki-stepper-step" +
                    (isActive ? " ki-stepper-step--active" : "") +
                    (isDone ? " ki-stepper-step--done" : "")
                    }
                    type="button"
                    onClick={() => onStepChange(index)}
                >
                    {!showLabels && (
                        <span className="ki-stepper-circle">
                        {isDone ? <Check size={18} strokeWidth={2}/> : index + 1}
                        </span>
                    )}
                    {showLabels && (
                        <span className="ki-stepper-label">{step.label}</span>
                    )}
                </button>
                );
            })}

            {hasAdd && (
                <button
                type="button"
                className="ki-stepper-add"
                onClick={onAddStep}
                title="Neuen Step hinzufügen"
                >
                +
                </button>
            )}
            </div>
        </div>
      </div>

      <div className="ki-stepper-body">
        {steps[currentIndex] ? steps[currentIndex].content : (
          <div className="ki-stepper-empty">
            Kein Step ausgewählt.
          </div>
        )}
      </div>
    </div>
  );
}
