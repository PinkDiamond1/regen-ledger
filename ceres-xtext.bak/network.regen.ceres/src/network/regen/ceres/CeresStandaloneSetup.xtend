/*
 * generated by Xtext 2.13.0
 */
package network.regen.ceres


/**
 * Initialization support for running Xtext languages without Equinox extension registry.
 */
class CeresStandaloneSetup extends CeresStandaloneSetupGenerated {

	def static void doSetup() {
		new CeresStandaloneSetup().createInjectorAndDoEMFRegistration()
	}
}