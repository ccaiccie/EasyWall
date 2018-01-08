module.exports = function (grunt) {
    grunt.initConfig({
        pkg: grunt.file.readJSON('package.json'),
        uglify: {
            my_target: {
                files: {
                    '../assets/javascript/bootstrap.min.js': ['./node_modules/bootstrap/dist/js/bootstrap.js'],
                    '../assets/javascript/jquery.min.js': ['./node_modules/jquery/dist/jquery.js'],
                    '../assets/javascript/popper.min.js': ['./node_modules/popper.js/dist/umd/popper.js'],
                    '../assets/javascript/easywall.min.js': ['./js/easywall.js']
                }
            },
            options: {
                'mangle': true,
                'report': 'gzip',
            }
        },
        cssmin: {
            target: {
                files: {
                    '../assets/stylesheet/bootstrap.min.css': ['./node_modules/bootstrap/dist/css/bootstrap.css'],
                    '../assets/stylesheet/font-awesome.min.css': ['./node_modules/font-awesome/css/font-awesome.css'],
                    '../assets/stylesheet/easywall.min.css': ['./css/easywall.css']
                }
            },
            options: {
                report: 'gzip',
                mergeIntoShorthands: false,
                roundingPrecision: -1
            }
        }
    });

    grunt.loadNpmTasks('grunt-contrib-uglify-es');
    grunt.loadNpmTasks('grunt-contrib-cssmin');

    grunt.registerTask('default', ['uglify', 'cssmin']);
};