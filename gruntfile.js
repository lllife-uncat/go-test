
module.exports = function(grunt) {
    grunt.initConfig({
        watch: {
            go: {
                files: ["test/**/*.go"],
                tasks: ["shell:test"]
            }
        },

        shell: {
            test: {
                command: "go test test/*.go -v"
            }
        }
    });

    grunt.loadNpmTasks("grunt-shell");
    grunt.loadNpmTasks("grunt-contrib-watch");
    grunt.registerTask("default", ["watch"]);
};