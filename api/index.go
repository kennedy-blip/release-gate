package api
import("encoding/json";"net/http";"strings")
func Handler(w http.ResponseWriter,r *http.Request){w.Header().Set("Cache-Control","no-store");path:=strings.Trim(r.URL.Query().Get("path"),"/");if path==""{path=strings.Trim(strings.TrimPrefix(r.URL.Path,"/api/"),"/")};switch path{case "releases":ReleasesHandler(w,r);case "evaluate":EvaluateHandler(w,r);default:writeError(w,404,"unknown API route")}}
func writeJSON(w http.ResponseWriter,status int,v any){w.Header().Set("Content-Type","application/json; charset=utf-8");w.WriteHeader(status);_ = json.NewEncoder(w).Encode(v)}
func writeError(w http.ResponseWriter,status int,message string){writeJSON(w,status,map[string]string{"error":message})}
