package main

if err = (&controllers.AWSRoleAssignmentReconciler{
	Client: mgr.GetClient(),
	Scheme: mgr.GetScheme(),
}).SetupWithManager(mgr); err != nil {
	log.Error(err, "unable to create controller", "controller", "AWSRoleAssignment")
	os.Exit(1)
}