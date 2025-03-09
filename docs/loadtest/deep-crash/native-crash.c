#include <node_api.h>
#include <stdlib.h>
#include <stdio.h>

// Function that will cause a segmentation fault
napi_value ForceCrash(napi_env env, napi_callback_info info) {
  printf("Forcing crash now...\n");
  fflush(stdout);
  
  // Access invalid memory - this will definitely crash the process
  int *p = NULL;
  *p = 42;  // This causes a segmentation fault
  
  // The following won't execute because of the crash
  napi_value result;
  napi_create_int32(env, 0, &result);
  return result;
}

// Initialize the module
napi_value Init(napi_env env, napi_value exports) {
  napi_value fn;
  
  // Create the crash function
  napi_create_function(env, NULL, 0, ForceCrash, NULL, &fn);
  napi_set_named_property(env, exports, "crash", fn);
  
  return exports;
}

// Register the module
NAPI_MODULE(NODE_GYP_MODULE_NAME, Init)
