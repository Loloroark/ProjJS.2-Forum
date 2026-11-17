func (a *App) Start() {
	go func() {
		if err := http.ListenAndServe(":8080", a.ApiRouter); err != nil {
			log.Fatalf("Erreur lancement API - %s", err.Error())
		}
	}()

	port := config.GetEnvWithDefault("PORT", "8081")
	if err := http.ListenAndServe(":"+port, a.WebRouter); err != nil {
		log.Fatalf("Erreur lancement serveur web - %s", err.Error())
	}
}